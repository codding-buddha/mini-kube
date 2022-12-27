package manager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/codding-buddha/mini-kube/common"
	"github.com/codding-buddha/mini-kube/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[uuid.UUID]*task.Task
	EventDb       map[uuid.UUID]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
	LastWorkerId  int
}

func (m *Manager) AddTask(te task.TaskEvent) {
	m.Pending.Enqueue(te)
}

func (m *Manager) GetTasks() []*task.Task {
	tasks := make([]*task.Task, len(m.TaskDb))
	for _, task := range m.TaskDb {
		tasks = append(tasks, task)
	}

	return tasks
}

func (m *Manager) SelectWorker() string {
	nextWorker := (m.LastWorkerId + 1) % len(m.Workers)
	m.LastWorkerId = nextWorker
	return m.Workers[nextWorker]
}

func (m *Manager) updateTasks() {
	for _, worker := range m.Workers {
		log.Printf("Checking worker %v for task updates.", worker)
		url := fmt.Sprintf("http://%s/tasks", worker)
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Error connecting to %v:%v", worker, err)
		}

		if resp.StatusCode != http.StatusOK {
			log.Printf("Error sending request: %v", err)
		}

		d := json.NewDecoder(resp.Body)
		var tasks []*task.Task
		err = d.Decode(&tasks)
		if err != nil {
			log.Printf("Error unmarshalling tasks: %s", err.Error())
		}

		for _, t := range tasks {
			log.Printf("Attempting to update task %v", t.ID)
			_, ok := m.TaskDb[t.ID]
			if !ok {
				log.Printf("Task with ID %s not found\n", t.ID)
				return
			}

			if m.TaskDb[t.ID].State != t.State {
				m.TaskDb[t.ID].State = t.State
			}

			m.TaskDb[t.ID].StartTime = t.StartTime
			m.TaskDb[t.ID].FinishTime = t.FinishTime
			m.TaskDb[t.ID].ContainerID = t.ContainerID
		}
	}
}

func (m *Manager) UpdateTasks() {
	for {
		fmt.Printf("[Manager] Updating tasks from %d workers\n.", len(m.Workers))
		m.updateTasks()
		time.Sleep(15 * time.Second)
	}
}

func (m *Manager) SendWork() {
	if m.Pending.Len() == 0 {
		log.Println("No work in the queue.")
		return
	}

	w := m.SelectWorker()
	e := m.Pending.Dequeue()
	te := e.(task.TaskEvent)
	t := te.Task
	log.Printf("Pulled %v off pending queue", t)

	m.EventDb[te.ID] = &te
	m.WorkerTaskMap[w] = append(m.WorkerTaskMap[w], te.Task.ID)
	m.TaskWorkerMap[t.ID] = w
	t.State = task.Scheduled
	m.TaskDb[t.ID] = &t
	data, err := json.Marshal(te)
	if err != nil {
		log.Printf("Unable to marshall object obj:%v, error:%v", t, err)
	}

	url := fmt.Sprintf("http://%s/tasks", w)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Error connecting %v:%v", w, err)
		m.Pending.Enqueue(t)
		return
	}

	d := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		e := common.ErrResponse{}
		err := d.Decode(&e)
		if err != nil {
			fmt.Printf("Error decoding response: %s\n", err)
			return
		}

		log.Printf("Response error (%d): %s", e.HTTPStatusCode, e.Message)
		return
	}

	t = task.Task{}
	err = d.Decode(&t)

	if err != nil {
		fmt.Printf("Error decoding response: %s\n", err.Error())
		return
	}
	log.Printf("%#v\n", t)
}

func (m *Manager) ProcessTasks() {
	for {
		log.Println("Processing any task in the queue")
		m.SendWork()
		log.Println("Sleeping for 10 seconds")
		time.Sleep(10 * time.Second)
	}
}

func New(workers []string) *Manager {
	taskDb := make(map[uuid.UUID]*task.Task)
	eventDb := make(map[uuid.UUID]*task.TaskEvent)
	workerTaskMap := make(map[string][]uuid.UUID)
	taskWorkerMap := make(map[uuid.UUID]string)
	for worker := range workers {
		workerTaskMap[workers[worker]] = []uuid.UUID{}
	}

	return &Manager{
		Pending:       *queue.New(),
		Workers:       workers,
		TaskDb:        taskDb,
		EventDb:       eventDb,
		WorkerTaskMap: workerTaskMap,
		TaskWorkerMap: taskWorkerMap,
	}
}

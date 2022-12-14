package worker

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/codding-buddha/mini-kube/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	Stats     Stats
	TaskCount int
}

func (w *Worker) InspectTask(t task.Task) task.DockerInspectResponse {
	config := task.NewConfig(&t)
	d := task.NewDocker(config)
	return d.Inspect(t.ContainerID)
}

func (w *Worker) AddTask(t task.Task) {
	w.Queue.Enqueue(t)
}

// GetTasks return all tasks of the worker.
func (w *Worker) GetTasks() []*task.Task {
	tasks := []*task.Task{}
	for _, task := range w.Db {
		tasks = append(tasks, task)
	}

	return tasks
}

func (w *Worker) StartTask(t task.Task) task.DockerResult {
	t.StartTime = time.Now().UTC()
	config := task.NewConfig(&t)
	d := task.NewDocker(config)
	result := d.Run()
	if result.Error != nil {
		log.Printf("Error running task %v:%v\n", t.ID, result.Error)
		t.State = task.Failed
		w.Db[t.ID] = &t
		return result
	}

	t.ContainerID = result.ContainerId
	t.State = task.Running
	w.Db[t.ID] = &t
	return result
}

func (w *Worker) StopTask(t task.Task) task.DockerResult {
	config := task.NewConfig(&t)
	d := task.NewDocker(config)
	result := d.Stop(t.ContainerID)

	if result.Error != nil {
		log.Printf("Error stopping container %v:%v", t.ContainerID, result.Error)
	}
	t.FinishTime = time.Now().UTC()
	t.State = task.Completed
	w.Db[t.ID] = &t
	log.Printf("Stopped and removed container %v for task %v", t.ContainerID, t.ID)

	return result
}

func (w *Worker) CollectStats() {
	for {
		log.Println("Collecting stats")
		w.Stats = *GetStats()
		w.TaskCount = w.Stats.TaskCount
		time.Sleep(15 * time.Second)
	}
}

func (w *Worker) UpdateTasks() {
	for {
		log.Println("Checking status of tasks.")
		w.updateTasks()
		log.Println("Task update completed.")
		log.Println("Sleeping for 15 seconds.")
		time.Sleep(15 * time.Second)
	}
}

func (w *Worker) updateTasks() {
	for id, t := range w.Db {
		if t.State == task.Running {
			resp := w.InspectTask(*t)
			if resp.Error != nil {
				fmt.Printf("ERROR: %v", resp.Error)
			}

			if resp.Container == nil {
				log.Printf("No container for running task %s", id)
				w.Db[id].State = task.Failed
			}

			if resp.Container.State.Status == "exited" {
				log.Printf("Container for task %s in non-running state %s", id, resp.Container.State.Status)
				w.Db[id].State = task.Failed
			}

			log.Printf("Running on port %v.\n", resp.Container.NetworkSettings.NetworkSettingsBase.Ports)

			w.Db[id].HostPorts = resp.Container.NetworkSettings.NetworkSettingsBase.Ports
		}
	}

}

func (w *Worker) RunTasks() {
	for {
		if w.Queue.Len() != 0 {
			result := w.runTask()
			if result.Error != nil {
				log.Printf("Error running tasks: %v\n", result.Error)
			}
		} else {
			log.Printf("No tasks to process currently, task queue is empty.\n")
		}
		log.Println("Sleeping for 10 seconds")
		time.Sleep(10 * time.Second)
	}
}

func (w *Worker) runTask() task.DockerResult {
	t := w.Queue.Dequeue()
	if t == nil {
		log.Printf("No tasks in the queue")
		return task.DockerResult{Error: nil}
	}

	taskQueued := t.(task.Task)
	taskPersisted := w.Db[taskQueued.ID]

	if taskPersisted == nil {
		taskPersisted = &taskQueued
		w.Db[taskPersisted.ID] = &taskQueued
	}

	var result task.DockerResult

	if task.ValidStateTransition(taskPersisted.State, taskQueued.State) {
		switch taskQueued.State {
		case task.Scheduled:
			if taskQueued.ContainerID != "" {
				result = w.StopTask(taskQueued)
				if result.Error != nil {
					log.Printf("Error stopping existing container %v.\n", result.Error)
				}
			}

			result = w.StartTask(taskQueued)
		case task.Completed:
			result = w.StopTask(taskQueued)
		default:
			result.Error = errors.New("Invalid State! ")

		}
	} else {
		err := fmt.Errorf("Invalid transistion from %v to %v", taskPersisted.State, taskQueued.State)
		result.Error = err
	}

	return result
}

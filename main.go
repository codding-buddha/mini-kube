package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/codding-buddha/mini-kube/task"
	"github.com/codding-buddha/mini-kube/worker"
	"github.com/docker/docker/client"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {
	host := os.Getenv("MINI_KUBE_HOST")
	fmt.Println(os.Getenv("MINI_KUBE_PORT"))
	port, err := strconv.Atoi(os.Getenv("MINI_KUBE_PORT"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Starting worker and API at %v:%v\n", host, port)

	w := worker.Worker{
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]*task.Task),
	}

	api := worker.Api{Address: host, Port: port, Worker: &w}
	go runTasks(&w)
	go w.CollectStats()
	api.Start()
}

func runTasks(w *worker.Worker) {
	for {
		if w.Queue.Len() != 0 {
			result := w.RunTasks()
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

func createContainer() (*task.Docker, *task.DockerResult) {
	c := task.Config{
		Name:  "test-container-1",
		Image: "postgres:13",
		Env: []string{
			"POSTGRES_USER=cube",
			"POSTGRES_PASSWORD=secret",
		},
	}

	dc, _ := client.NewClientWithOpts(client.FromEnv)
	d := task.Docker{
		Config: c,
		Client: dc,
	}

	result := d.Run()
	if result.Error != nil {
		fmt.Printf("%v\n", result.Error)
		return nil, nil
	}

	fmt.Printf("Container %s is running with config %v\n", result.ContainerId, c)
	return &d, &result
}

func stopContainer(d *task.Docker, id string) *task.DockerResult {
	result := d.Stop(id)
	if result.Error != nil {
		fmt.Printf("%v\n", result.Error)
		return nil
	}

	fmt.Printf("Container %s has been stopped and removed\n", result.ContainerId)
	return &result
}

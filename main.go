package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codding-buddha/mini-kube/manager"
	"github.com/codding-buddha/mini-kube/task"
	"github.com/codding-buddha/mini-kube/worker"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {
	whost := os.Getenv("MINI_KUBE_WORKER_HOST")
	wport, err := strconv.Atoi(os.Getenv("MINI_KUBE_WORKER_PORT"))
	if err != nil {
		panic(err)
	}

	mhost := os.Getenv("MINI_KUBE_MANAGER_HOST")
	mport, err := strconv.Atoi(os.Getenv("MINI_KUBE_MANAGER_PORT"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Starting worker and API at %v:%v\n", whost, wport)

	w := worker.Worker{
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]*task.Task),
	}

	wapi := worker.Api{Address: whost, Port: wport, Worker: &w}
	go w.RunTasks()
	go w.CollectStats()
	go wapi.Start()

	workers := []string{fmt.Sprintf("%s:%d", whost, wport)}
	fmt.Printf("Starting manager and API at %v:%v\n", mhost, mport)
	m := manager.New(workers)
	mapi := manager.Api{Address: mhost, Port: mport, Manager: m}
	go m.ProcessTasks()
	go m.UpdateTasks()

	mapi.Start()
}

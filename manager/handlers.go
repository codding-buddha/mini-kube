package manager

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/codding-buddha/mini-kube/common"
	"github.com/codding-buddha/mini-kube/task"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (api *Api) StartTaskHandler(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	te := task.TaskEvent{}
	err := d.Decode(&te)

	if err != nil {
		msg := fmt.Sprintf("Error unmarshalling body: %v\n", err)
		log.Printf(msg)
		w.WriteHeader(http.StatusBadRequest)

		e := common.ErrResponse{
			HTTPStatusCode: http.StatusBadRequest,
			Message:        msg,
		}

		json.NewEncoder(w).Encode(e)
		return
	}

	api.Manager.AddTask(te)
	log.Printf("Added task %v\n", te.Task.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(te.Task)
}

func (api *Api) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(api.Manager.GetTasks())
}

func (api *Api) StopTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "taskID")

	if taskID == "" {
		log.Printf("No taskID passed in request.\n")
		w.WriteHeader(http.StatusBadRequest)
	}

	tID, _ := uuid.Parse(taskID)
	_, ok := api.Manager.TaskDb[tID]
	if !ok {
		log.Printf("No task with ID %v found.", tID)
		w.WriteHeader(http.StatusBadRequest)
	}

	te := task.TaskEvent{
		ID:        uuid.New(),
		State:     task.Completed,
		Timestamp: time.Now(),
	}

	taskToStop := api.Manager.TaskDb[tID]
	taskCopy := *taskToStop
	taskCopy.State = task.Completed
	te.Task = *taskToStop
	api.Manager.AddTask(te)
	log.Printf("Added task event %v, to stop task %v\n", te.ID, taskToStop.ID)
	w.WriteHeader(http.StatusAccepted)
}

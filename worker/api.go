package worker

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ErrResponse struct {
	HTTPStatusCode int
	Message        string
}

type Api struct {
	Address string
	Port    int
	Worker  *Worker
	Router  *chi.Mux
}

func (api *Api) initRouter() {
	api.Router = chi.NewRouter()
	api.Router.Route("/stats", func(r chi.Router) {
		r.Get("/", api.GetStatsHandler)
	})
	api.Router.Route("/tasks", func(r chi.Router) {
		r.Post("/", api.StartTaskHandler)
		r.Get("/", api.GetTasksHandler)
		r.Route("/{taskID}", func(r chi.Router) {
			r.Delete("/", api.StopTaskHandler)
		})
	})
}

func (api *Api) Start() {
	api.initRouter()
	http.ListenAndServe(fmt.Sprintf("%s:%d", api.Address, api.Port), api.Router)
}

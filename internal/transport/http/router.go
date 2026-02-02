package router

import (
	"net/http"
	handler "newTaskManagerApi/internal/transport/http/v1/task"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewRouter(taskHandler *handler.TaskHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			"GET", "PATCH", "POST", "DELETE",
		},
	}))

	r.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
	})

	r.Route(
		"/api", func(r chi.Router) {
			r.Route("/v1", func(r chi.Router) {
				r.Route("/todos", func(r chi.Router) {
					r.Post("/", taskHandler.Create)
					r.Get("/", taskHandler.GetAllTasks)
					r.Patch("/{id}", taskHandler.Update)
					r.Delete("/{id}", taskHandler.Delete)
				})
			})
		},
	)

	return r
}

package handler

import (
	"encoding/json"
	"net/http"
	"newTaskManagerApi/internal/domain"
	"newTaskManagerApi/internal/service"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

type TaskRead struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TaskCreate struct {
	Title string `json:"title"`
}

type TaskUpdateStatus struct {
	Completed bool `json:"completed"`
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req TaskCreate

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondError(w, statusCode(err), "invalid body")
		return
	}

	task, err := h.service.Create(r.Context(), req.Title)
	if err != nil {
		respondError(w, statusCode(err), err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, toDTO(*task))
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAllTasks(r.Context())
	if err != nil {
		respondError(w, statusCode(err), err.Error())
		return
	}

	DTOTasks := []TaskRead{}
	for _, v := range tasks {
		DTOTasks = append(DTOTasks, toDTO(*v))
	}

	respondJSON(w, statusCode(err), DTOTasks)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req TaskUpdateStatus
	id := chi.URLParam(r, "id")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondError(w, statusCode(err), "invalid body")
		return
	}

	task, err := h.service.UpdateTaskStatus(r.Context(), id, req.Completed)
	if err != nil {
		respondError(w, statusCode(err), err.Error())
		return
	}

	respondJSON(w, statusCode(err), toDTO(*task))
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.DeleteTask(r.Context(), id)
	if err != nil {
		respondError(w, statusCode(err), err.Error())
		return
	}

	respondJSON(w, statusCode(err), map[string]string{"message": "Deleted"})
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

func toDTO(domainModel domain.Task) TaskRead {
	return TaskRead{
		ID:        domainModel.ID,
		Title:     domainModel.Title,
		Completed: domainModel.Completed,
	}
}

func statusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	if err == domain.ErrNotFound {
		return http.StatusNotFound
	}
	if err == domain.ErrAlreadyExists {
		return http.StatusConflict
	}
	if err == domain.ErrInvalidTitle {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

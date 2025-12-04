package controller

import (
	"KageNoEn/model"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (c *Controller) GetAllUserStatus(w http.ResponseWriter, r *http.Request) {
	userStatus, err := c.R.GetAllUserStatus()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := &model.UserStatusList{
		Data: userStatus,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (c *Controller) CreateUserStatus(w http.ResponseWriter, r *http.Request) {
	var status model.UserStatus
	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUUID := uuid.New().String()
	status.Id = newUUID

	if err := c.R.CreateUserStatus(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := &model.UserStatusResponse{
		Data: status,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func (c *Controller) UpdateUserStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	status, err := c.R.GetUserStatusById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.R.UpdateUserStatus(&status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := &model.UserStatusResponse{
		Data: status,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (c *Controller) DeleteUserStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := c.R.DeleteUserStatus(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := &model.DeleteModel{
		Message: "success fully deleted",
		Status: http.StatusNoContent,
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(res)
}

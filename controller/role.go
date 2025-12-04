package controller

import (
	"KageNoEn/model"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (c *Controller) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := c.R.GetAllRoles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roles)
}

func (c *Controller) GetRole(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	role, err := c.R.GetRole(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(role)
}

func (c *Controller) CreateRole(w http.ResponseWriter, r *http.Request) {
	var role model.Role
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUUID := uuid.New().String()
	role.Id = newUUID

	if err := c.R.CreateRole(role); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(role)
}

func (c *Controller) DeleteRole(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := c.R.DeleteRole(id); err != nil {
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

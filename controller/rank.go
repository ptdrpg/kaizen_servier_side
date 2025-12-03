package controller

import (
	"KageNoEn/model"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type RankList struct {
	Data []model.Rank `json:"data"`
}

type RankResponse struct {
	Data model.Rank `json:"data"`
}

func (c *Controller) GetAllRanks(w http.ResponseWriter, r *http.Request) {
	ranks, err := c.R.GetAllRanks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := RankList{Data: ranks}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (c *Controller) CreateRank(w http.ResponseWriter, r *http.Request) {
	var rank model.Rank
	err := json.NewDecoder(r.Body).Decode(&rank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUUID := uuid.New().String()
	rank.Id = newUUID

	if err := c.R.CreateRank(rank); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := RankResponse{
		Data: rank,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func (c *Controller) UpdateRank(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	rank, err := c.R.GetRankById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&rank); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.R.UpdateRank(&rank); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := RankResponse{
		Data: rank,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (c *Controller) DeleteRank(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := c.R.DeleteRank(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("success fully deleted")
}

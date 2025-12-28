package controller

import (
	"KageNoEn/model"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (c *Controller) GetAllFriends(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	friends, err := c.R.GetAllFriends(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := model.FriendListResponse{
		Data: friends,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (c *Controller) GetRequest(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	requests, err := c.R.GetFriendRequest(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := model.FriendRequestList{
		Data: requests,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (c *Controller) AddFriend(w http.ResponseWriter, r *http.Request) {
	var input model.AddFriendType
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	invit := model.FriendList{
		Id: uuid.New().String(),
		Sender: input.SenderId,
		Receiver: input.ReceiverId,
		Status: "pending",
	}

	if err := c.R.AddFriend(invit); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := model.UpdateMessage{
		Message: "Friend request sent successfully",
		Status: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (c *Controller) ConfirmFriend(w http.ResponseWriter, r *http.Request) {
	invitId := chi.URLParam(r, "id")
	invitation, err := c.R.GetInvitationByID(invitId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if invitation.Status != "pending" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	invitation.Status = "accepted"

	if err := c.R.ConfirmFriend(invitation); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := model.UpdateMessage{
		Message: "Friend request confirmed successfully",
		Status: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json") 	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
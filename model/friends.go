package model

import "time"

type FriendList struct {
	Id        string    `gorm:"id;primarykey" json:"id"`
	Sender    string    `gorm:"sender" json:"sender"`
	Receiver  string    `gorm:"receiver" json:"receiver"`
	Status    string    `gorm:"status" json:"status"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

type FriendRequestType struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type FriendRequestList struct {
	Data []FriendRequestType `json:"data"`
}

type FriendWithStatus struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Status   string `json:"status"`
	IsOnline bool   `json:"is_online"`
}

type FriendListResponse struct {
	Data []FriendWithStatus `json:"data"`
}

type AddFriendType struct {
	ReceiverId string `json:"receiver_id"`
	SenderId   string `json:"sender_id"`
}

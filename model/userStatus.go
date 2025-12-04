package model

import "time"

type UserStatus struct {
	Id        string    `gorm:"id;primaryKey" json:"id"`
	Label     string    `gorm:"label" json:"label"`
	IsOnline  bool      `gorm:"is_online" json:"is_online"`
	LastLogin time.Time `gorm:"last_login" json:"last_login"`
}

type UserStatusList struct {
	Data []UserStatus `json:"data"`
}

type UserStatusResponse struct {
	Data UserStatus `json:"data"`
}

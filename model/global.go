package model

import "time"

type DeleteModel struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type UpdateMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type SecureUserRes struct {
	Id        string      `gorm:"id;primarykey" json:"id"`
	Username  string      `gorm:"username;unique" json:"username"`
	RankId    string      `gorm:"rank_id" json:"rank_id"`
	Rank      *Rank       `gorm:"foreignKey:RankId;references:Id" json:"rank"`
	StatusId  string      `gorm:"status_id" json:"status_id"`
	Status    *UserStatus `gorm:"foreignKey:StatusId;references:Id" json:"status"`
	RoleId    string      `gorm:"role_id;" json:"role_id"`
	Role      *Role       `gorm:"foreignKey:RoleId;references:Id" json:"role"`
	MaxRank   string      `gorm:"max_rank" json:"max_rank"`
	Elo       int         `gorm:"elo;default:0" json:"elo"`
	IsOnline  bool        `gorm:"is_online;default:false" json:"is_online"`
	LastLogin time.Time   `gorm:"last_login;default:null" json:"last_login"`
	CreatedAt time.Time   `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time   `gorm:"updated_at" json:"updated_at"`
}

type RegisterResponse struct {
	Data  SecureUserRes `json:"data"`
	Token string        `json:"token"`
}

type ChangePassInput struct {
	OldPass string `json:"old_pass"`
	NewPass string `json:"new_pass"`
}

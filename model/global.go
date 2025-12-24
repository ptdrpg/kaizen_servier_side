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
	Id        string      `json:"id"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	RankId    string      `json:"rank_id"`
	Rank      *Rank       `json:"rank"`
	StatusId  string      `json:"status_id"`
	Status    *UserStatus `json:"status"`
	RoleId    string      `json:"role_id"`
	Role      *Role       `json:"role"`
	MaxRank   string      `json:"max_rank"`
	Elo       int         `json:"elo"`
	IsOnline  bool        `json:"is_online"`
	LastLogin time.Time   `json:"last_login"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type RegisterResponse struct {
	Data  SecureUserRes `json:"data"`
	Token string        `json:"token"`
}

type ChangePassInput struct {
	OldPass string `json:"old_pass"`
	NewPass string `json:"new_pass"`
}

package model

import "time"

type User struct {
	Id        string    `gorm:"id;primarykey" json:"id"`
	Username  string    `gorm:"username;unique" json:"username"`
	Email     string    `gorm:"email;unique" json:"email"`
	Password  string    `gorm:"password" json:"password"`
	RankId    string    `gorm:"rank_id" json:"rank_id"`
	RoleId    string    `gorm:"role_id" json:"role_id"`
	MaxRank   string    `gorm:"max_rank" json:"max_rank"`
	Elo       int       `gorm:"elo" json:"elo"`
	GuildId   string    `gorm:"guild_id;default:null" json:"guild_id"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

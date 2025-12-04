package model

type Role struct {
	Id    string `gorm:"id;primarykey" json:"id"`
	Label string `gorm:"label" json:"label"`
}

package entity

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

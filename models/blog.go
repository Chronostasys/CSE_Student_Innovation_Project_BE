package models

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	Team_Id    uint
	Auth_Email string
	Title      string
	Content    string
}
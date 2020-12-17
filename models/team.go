package models

import "github.com/jinzhu/gorm"

type Team struct {
	gorm.Model
	Name          string
	Owner_email   string
	Description   string
	Mile_Stone string
}
type TeamMember struct {
	gorm.Model
	Team_Id uint
	Email   string
}
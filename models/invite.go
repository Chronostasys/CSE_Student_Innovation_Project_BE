package models

import "github.com/jinzhu/gorm"

type Invite struct {
	gorm.Model
	Owner_email   	string
	Team_Id	  		int
	MemberInvited 	string
	Is_accepted     int
}

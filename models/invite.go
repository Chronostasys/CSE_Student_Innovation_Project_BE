package models

import "github.com/jinzhu/gorm"

type Invite_Record struct {
	gorm.Model
	Team_Id	  		int
	MemberInvited 	string
	Is_accepted     int
}

package services

import "github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"

func AddTeamMemberInDB(invite models.Invite){
	db.Create(&invite)
}
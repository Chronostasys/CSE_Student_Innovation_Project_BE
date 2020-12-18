package services

import "github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"

func AddInvite_RecordInDB(a models.Invite_Record){
	db.Create(&a)
}

package services

import "github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"

func AddTeamMemberInDB(a models.TeamMember){
	db.Create(&a)
}
func AddInvite_RecordInDB(a models.Invite_Record){
	db.Create(&a)
}
func IsUrlReal(team_id int,email_accept string)(IsUrlReal bool){
	var Invite_Record models.Invite_Record
	err1:=db.Where("Team_Id=?", team_id).Find(&Invite_Record).Error
	if err1!=nil{
		return false
	}
	err2:=db.Where("MemberInvited=?", email_accept).Find(&Invite_Record).Error
	if err2!=nil{
		return false
	}
	err3:=db.Where("Is_accepted=?", 0).Find(&Invite_Record).Error
	if err3!=nil{
		return false
	}
	return true
}

package services

import "github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"

func AddTeamMemberInDB(a models.TeamMember){
	db.Create(&a)
}
func IsUrlReal(team_id int,email_accept string)(IsUrlReal bool){
	var Invite_Record models.Invite_Record
	err:=db.Where("Team_Id=? and MemberInvited=? and Is_accepted=? ", team_id,email_accept,0).Find(&Invite_Record)
	if err!=nil{
		return false
	}
	return true
}


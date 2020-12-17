package services

import (
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
)

func AddTeamInDB(team models.Team){
		db.Create(&team)
}
func DeleteTeamInDB(teamName string){
	var team models.Team
	db.Where("Name=?", teamName).Find(&team)
	db.Delete(team)
	/*如果删除组织的话，所有属于这个组织的人都会自动退出这个组织*/
	/*这里需要调用退出组织的方法*/
}
func IsTeamExist(teamName string)(is_team_exist bool){
	var team models.Team
	err:=db.Where("Name=?", teamName).Find(&team).Error
	if err!=nil{
		return false
	}
	return true
}
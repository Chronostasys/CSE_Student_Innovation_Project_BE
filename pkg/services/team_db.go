package services

import (
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
)

func AddTeamInDB(team models.Team){
		db.Create(&team)
}
func DeleteTeamInDB(teamID uint){
	var team models.Team
	db.Where("ID=?", teamID).Find(&team)
	db.Delete(team)
	/*如果删除组织的话，所有属于这个组织的人都会自动退出这个组织*/
	/*这里需要调用退出组织的方法*/
}
func IsTeamExist(teamID uint)(is_team_exist bool){
	var team models.Team
	err:=db.Where("ID=?", teamID).Find(&team).Error
	if err!=nil{
		return false
	}
	return true
}
func ChangeTeamDetailInDB(teamID uint,newTeamName string,mileStoneUpdate string,descriptionUpdate string){
	var team models.Team
	db.Where("ID=?", teamID).Find(&team)

	team.Description = descriptionUpdate
	team.Name = newTeamName
	team.Mile_Stone = mileStoneUpdate

	db.Save(&team)

}
func GetTeams(startID uint,listSize int)(teams []models.Team){
	db.Order("ID").Offset(startID).Limit(listSize).Find(&teams)
	return
}
package services

import (
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
)

func AddTeamInDB(email string,teamName string,description string,mile_stone string){
		team:=models.Team{Owner_email:email,Name:teamName,Description:description,Mile_Stone:mile_stone}
		db.Create(&team)
}
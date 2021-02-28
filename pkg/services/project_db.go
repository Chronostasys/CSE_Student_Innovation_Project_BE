package services

import "github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"

func GetProjects(page int,listSize int,isDescent bool)(projects []models.Project){
	if isDescent{
		db.Offset(page* listSize).Limit(listSize).Order("created_at desc").Find(&projects)
	}else{
		db.Offset(page* listSize).Limit(listSize).Find(&projects)
	}
	return projects
}

func GetOneProject(projectId int)(project models.Project,err error){
	err=db.Where("id=?",projectId).Find(&project).Error
	return
}
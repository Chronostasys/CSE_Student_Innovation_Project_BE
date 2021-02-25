package services

import "github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"

func CreateEnterprise(enterprise models.Enterprise)(err error){
	err=db.Create(enterprise).Error
	return err
}

func CreateProject(project models.Project)(err error){
	err=db.Create(project).Error
	return err
}
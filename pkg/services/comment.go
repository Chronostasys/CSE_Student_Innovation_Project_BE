package services

import "github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"

func AddComment(comment *models.Comment)(err error){
	err=db.Create(comment).Error
	return err
}
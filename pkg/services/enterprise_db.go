package services

import "github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"

func GetEnterprises(page int,listSize int,isDescent bool)(enterprises []models.Enterprise){
	if isDescent{
		db.Offset(page* listSize).Limit(listSize).Order("created_at desc").Find(&enterprises)
	}else{
		db.Offset(page* listSize).Limit(listSize).Find(&enterprises)
	}
	return enterprises
}

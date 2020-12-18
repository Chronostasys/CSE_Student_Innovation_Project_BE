package services

import "github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"

func GetTeamMemberFromEmail(auth_email string)(teamMember models.TeamMember){
	db.Where("email = ?",auth_email).First(&teamMember)
	return teamMember
}

func AddBlog(auth_email string,title string,content string,team_id uint)(blog models.Blog){
	blog = models.Blog{
		Team_Id:team_id,
		Auth_Email: auth_email,
		Title: title,
		Content: content,
	}
	db.Create(&blog)
	return blog
}

func DeleteBlog(id int,auth_email string)(isDeleted bool,err error){
	var blog models.Blog
	db.Where("id = ?",id).First(&blog)
	if blog.Auth_Email!=auth_email{
		isDeleted=false
		return
	}
	err=db.Where("id = ?",id).Delete(&blog).Error
	if err!=nil{
		isDeleted=false
		return
	}
	return true,nil
}
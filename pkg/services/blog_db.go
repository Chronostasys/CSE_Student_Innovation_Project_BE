package services

import "github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"

func GetTeamMemberFromEmail(auth_email string)(teamMember models.TeamMember){
	db.Where("email = ?",auth_email).First(&teamMember)
	return teamMember
}

func AddBlog(blog models.Blog)(err error){
	err=db.Create(&blog).Error
	return err
}

func DeleteBlog(blog_id uint,auth_email string)(isDeleted bool,err error){
	var blog models.Blog
	db.Where("blog_id = ?",blog_id).First(&blog)
	if blog.Auth_Email!=auth_email{
		isDeleted=false
		return
	}
	err=db.Where("blog_id = ?",blog_id).Delete(&blog).Error
	if err!=nil{
		isDeleted=false
		return
	}
	return true,nil
}

func GetBlogs(page int,list_size int,is_descend bool)(blogs []models.Blog){
	if is_descend {
		db.Offset(page*list_size).Limit(list_size).Order("created_at desc").Find(&blogs)
	} else {
		db.Offset(page*list_size).Limit(list_size).Find(&blogs)
	}
	return blogs
}
func GetBlogsNumber()(number int){
	var blog models.Blog
	db.Model(&blog).Count(&number)
	return number
}
func GetOneBlog(blog_id int)(blog models.Blog,err error){
	err=db.Where("id=?",blog_id).Find(blog_id).Error
	return
}
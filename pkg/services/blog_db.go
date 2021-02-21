package services

import "github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"

//func GetTeamMemberFromEmail(auth_email string)(teamMember models.TeamMember){
//	db.Where("email = ?",auth_email).First(&teamMember)
//	return teamMember
//}

func AddBlog(blog *models.Blog)(err error){
	err=db.Create(blog).Error
	return err
}

func DeleteBlog(blogId uint,authEmail string)(isDeleted bool,err error){
	var blog models.Blog
	db.Where("blog_id = ?",blogId).First(&blog)
	if blog.AuthEmail!=authEmail{
		isDeleted=false
		return
	}
	err=db.Where("blog_id = ?",blogId).Delete(&blog).Error
	if err!=nil{
		isDeleted=false
		return
	}
	return true,nil
}


func GetBlogs(page int, listSize int, isDescend bool)(blogs []models.Blog){
	if isDescend {
		db.Offset(page* listSize).Limit(listSize).Order("created_at desc").Find(&blogs)
	} else {
		db.Offset(page* listSize).Limit(listSize).Find(&blogs)
	}
	return blogs
}
func GetBlogsNumber()(number int){
	var blog models.Blog
	db.Model(&blog).Count(&number)
	return number
}
func GetOneBlog(blogId int)(blog models.Blog,err error){
	err=db.Where("id=?", blogId).Find(&blog).Error
	return
}
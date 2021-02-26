package router

import (
	"fmt"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/services"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddBlog(context *gin.Context) {
	title:=context.PostForm("title")
	content:=context.PostForm("content")
	auth_email,_:=util.GetEmailFromToken(context)
	if content==""||title==""{
		context.JSON(http.StatusBadRequest,gin.H{
			"msg":"请输入内容",
		})
		return
	}
	blog:=models.Blog{
		//Team_Id: team_id,
		Auth_Email: auth_email,
		Title: title,
		Content: content,
	}
	err:=services.AddBlog(blog)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{
			"msg":"发布文章失败",
		})
		return
	}
	context.JSON(http.StatusOK,gin.H{
			"blog_id":blog.ID,
	})
	return
}

//只能删除自己发布的文章
func DeleteBlog(context *gin.Context){
	blog_id,_:=strconv.Atoi(context.PostForm("blog_id"))
	auth_email,_:=util.GetEmailFromToken(context)
	isDeleted,err:=services.DeleteBlog(uint(blog_id),auth_email)
	if isDeleted==false&&err==nil{
		context.JSON(http.StatusBadRequest,gin.H{
			"msg":"无权限删除文章",
		})
		return
	}
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{
			"msg":"删除失败，请重试",
		})
		return
	}
	context.JSON(http.StatusOK,gin.H{
		"msg":"删除成功",
	})
	return
}

func GetBlogs(context *gin.Context){
	page,_:=strconv.Atoi(context.Query("page"))
	list_size,_:=strconv.Atoi( context.Query("blog_num"))
	fmt.Println(page,list_size)
	blogs:=services.GetBlogs(page,list_size,true)
	fmt.Println(blogs)
	results:=[]map[string]interface{}{}
	for _,temp:=range blogs{
		user,_:=services.GetUserByEmail(temp.Auth_Email)
		time,_:=util.ConvertShanghaiTimeZone(temp.CreatedAt)
		results=append(results,map[string]interface{}{
			"blog_id":temp.ID,
			"title":temp.Title,
			"content":temp.Content,
			"auth_email":temp.Auth_Email,
			"author_name":user.Username,
			"publish_time":time.String(),
		})
	}
	context.JSON(http.StatusOK,gin.H{
		"msg":results,
	})
}


func GetBlogsNumber(context *gin.Context){

	context.JSON(http.StatusOK,gin.H{
		"total_number":services.GetBlogsNumber(),
	})
}
func GetBlog(context *gin.Context){
	blog_id,_:=strconv.Atoi(context.Param("blog_id"))
	blog,err:=services.GetOneBlog(blog_id)
	if err!=nil{
		context.JSON(404,gin.H{`err`:err.Error()})
	}else {
		user,_:=services.GetUserByEmail(blog.Auth_Email)
		time,_:=util.ConvertShanghaiTimeZone(blog.CreatedAt)
		context.JSON(200,gin.H{
			"blog_id":      blog.ID,
			"title":        blog.Title,
			"content":      blog.Content,
			"auth_email":   blog.Auth_Email,
			"author_name":  user.Username,
			"publish_time": time.String(),
		})


		return
	}
}

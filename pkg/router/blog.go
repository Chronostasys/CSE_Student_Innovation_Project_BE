package router

import (
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
	teamId,_:=strconv.Atoi(context.PostForm("team_id"))
	team_id:=uint(teamId)
	auth_email,_:=util.GetEmailFromCookie(context)
	teamMember:=services.GetTeamMemberFromEmail(auth_email)
	if teamMember.ID!=team_id{
		context.JSON(http.StatusBadRequest,gin.H{
			"msg":"不在指定队伍中，无法发布",
		})
		return
	}
	if content==""||title==""{
		context.JSON(http.StatusBadRequest,gin.H{
			"msg":"请输入内容",
		})
		return
	}
	blog:=models.Blog{
		Team_Id: team_id,
		Auth_Email: auth_email,
		Title: title,
		Content: content,
	}
	err:=services.AddBlog(blog)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{
			"msg":"发布文章失败",
		})
	}
	context.JSON(http.StatusOK,gin.H{
			"msg":"发布文章成功",
	})
}

//只能删除自己发布的文章
func DeleteBlog(context *gin.Context){
	blog_id,_:=strconv.Atoi(context.PostForm("blog_id"))
	auth_email,_:=util.GetEmailFromCookie(context)
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
}

func GetBlogs(context *gin.Context){
	team_id,_:=strconv.Atoi(context.PostForm("team_id"))
	start_id,_:=strconv.Atoi(context.PostForm("start_id"))
	list_size,_:=strconv.Atoi( context.PostForm("list_size"))
	is_descend,_:=strconv.ParseBool(context.PostForm("is_descend"))
	auth_email,_:=util.GetEmailFromCookie(context)
	teamMember:=services.GetTeamMemberFromEmail(auth_email)
	if teamMember.Team_Id!=uint(team_id){
		context.JSON(http.StatusBadRequest,gin.H{
			"msg":"不在队伍中，无法查看",
		})
		return
	}
	if start_id<0||list_size<0{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "不合规的id或表长",
		})
		return
	}
	blogs:=services.GetBlogs(uint(team_id),start_id,list_size,is_descend)
	results:=[]map[string]interface{}{}
	for _,temp:=range blogs{
		results=append(results,map[string]interface{}{
			"title":temp.Title,
			"auth_email":temp.Auth_Email,
			"content":temp.Content,
		})
	}
	context.JSON(http.StatusOK,gin.H{
		"msg":results,
	})
}
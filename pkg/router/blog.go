package router

import (
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/services"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddBlog(context *gin.Context) {
	title:=context.PostForm("title")
	content:=context.PostForm("content")
	auth_email,_:=util.GetEmailFromCookie(context)
	teamMember:=services.GetTeamMemberFromEmail(auth_email)
	if teamMember.ID<0{
		context.JSON(http.StatusBadRequest,gin.H{
			"msg":"请先加入队伍",
		})
		return
	}
	if title==""{
		context.JSON(http.StatusBadRequest,gin.H{
			"msg":"请输入标题",
		})
		return
	}
	if content==""{
		context.JSON(http.StatusBadRequest,gin.H{
			"msg":"请输入内容",
		})
		return
	}
	_=services.AddBlog(auth_email,title,content,teamMember.Team_Id)
	context.JSON(http.StatusOK,gin.H{
			"msg":"发布文章成功",
	})
}
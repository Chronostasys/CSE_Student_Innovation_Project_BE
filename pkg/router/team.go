package router

import (
	"github.com/Pivot-Studio/Authorization-Template/pkg/util"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatTeam(context *gin.Context) {
	teamName:= context.PostForm("team_name")
	email,_:= util.GetEmailFromCookie(context)
	description:= context.PostForm("description")
	mile_stone:= context.PostForm("email")
	if teamName==""{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":  "组织名称不能为空",
		})
		return

	}
	if description==""{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":  "组织描述不能为空",
		})
		return

	}
	if mile_stone==""{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":  "项目进度不能为空",
		})
		return

	}
	services.AddTeamInDB(email,teamName,description,mile_stone)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "组织创建成功",
	})
	return

}
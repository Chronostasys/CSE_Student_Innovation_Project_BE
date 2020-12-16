package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatTeam(context *gin.Context) {
	teamName:= context.PostForm("team_name")
	email:= context.PostForm("email")
	description:= context.PostForm("description")
	mile_stone:= context.PostForm("mile_stone")
	if teamName==""{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":  "团队名称不能为空",
		})
		return

	}
	if description==""{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":  "团队描述不能为空",
		})
		return

	}
	if mile_stone==""{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":  "项目进度不能为空",
		})
		return

	}

	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "helloworld",
	})
	return

}
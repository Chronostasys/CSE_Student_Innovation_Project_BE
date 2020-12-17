package router

import (
	"github.com/Pivot-Studio/Authorization-Template/pkg/util"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatTeam(context *gin.Context) {
	teamName:= context.PostForm("team_name")
	email,_:= util.GetEmailFromCookie(context)
	description:= context.PostForm("description")
	mile_stone:= context.PostForm("email")
	team:=models.Team{
		Owner_email:email,
		Name:teamName,
		Description:description,
		Mile_Stone:mile_stone,
	}
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
	services.AddTeamInDB(team)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "组织创建成功",
	})
	return

}

func DeleteTeam(context *gin.Context) {
	teamName := context.PostForm("team_name")
	is_team_exist := services.IsTeamExist(teamName)
	services.DeleteTeamInDB(teamName)
	if !is_team_exist {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "该组织不存在",
		})
		return
	}else{
		context.JSON(http.StatusOK, gin.H{
			"msg": "组织删除成功",
		})
		return
	}
}
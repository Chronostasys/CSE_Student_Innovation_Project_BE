package router

import (
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AcceptInvite(context *gin.Context) {
	Owner_email := context.Query("Owner_email")
	team_id, _ := strconv.Atoi(context.Query("team_id"))
	email_accept := context.Query("email_accept")
	TeamMember := models.TeamMember{
		Team_Id: uint(team_id),
		Email:   email_accept,
	}
	if Owner_email == "" || team_id == 0 || email_accept == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "错误",
		})
		return
	}
	IsUrlReal := services.IsUrlReal(team_id, email_accept)
	if !IsUrlReal {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "URL无效",
		})
		return
	} else {
		services.AddTeamMemberInDB(TeamMember)
		context.JSON(http.StatusOK, gin.H{
			"msg": "您已成功加入组织",
		})
		//这里还要把is_accepted修改成1
		return
	}
}
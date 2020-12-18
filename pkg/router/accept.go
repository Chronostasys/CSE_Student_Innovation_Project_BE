package router

import (
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AcceptInvite(context *gin.Context){
	Owner_email:=context.Query("Owner_email")
	team_id,_ :=strconv.Atoi(context.Query("team_id"))
	email_accept:=context.Query("email_accept")
	invite:=models.Invite{
		Owner_email:Owner_email,
		Team_Id:team_id,
		MemberInvited:email_accept,
		Is_accepted:1,
	}
	if Owner_email==""{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":  "错误",
		})
		return
	}
	if team_id==0{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":  "组织不存在",
		})
		return
	}
	if email_accept==""{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":  "错误",
		})
		return
	}
	services.AddTeamMemberInDB(invite)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "您已成功加入组织",
	})
	return
}

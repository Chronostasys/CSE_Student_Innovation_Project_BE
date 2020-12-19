package router

import (
	"github.com/Pivot-Studio/Authorization-Template/pkg/ATutil"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func InviteTeamMember(context *gin.Context){
	email,_:= ATutil.GetEmailFromCookie(context)
	team_id,_ :=strconv.Atoi(context.PostForm("team_id"))
	description:=context.PostForm("description")
	memberInvited:=context.PostForm("email_accept")
	Invite_Record:=models.Invite_Record{
		Team_Id:team_id,
		MemberInvited:memberInvited,
		Is_accepted:0,
	}
	if team_id==0{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请添加组织id",
		})
		return
	}
	if description==""{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请备注邀请信息",
		})
		return
	}
	if memberInvited==""{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请添加需要邀请的用户的电子邮箱",
		})
		return
	}
	services.InviteMemberToTeam(email,team_id,description,memberInvited)
	services.AddInvite_RecordInDB(Invite_Record)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "已邀请",
	})
	return
}

//这个以后再写吧

//func RemoveTeamMember(context *gin.Context){
//	teamName := context.PostForm("team_name")
//	is_team_exist := services.IsTeamMemberExist(teamName)
//	services.DeleteTeamInDB(teamName)
//	if !is_team_exist {
//		context.JSON(http.StatusBadRequest, gin.H{
//			"msg": "该组织不存在",
//		})
//		return
//	}else{
//		context.JSON(http.StatusOK, gin.H{
//			"msg": "组织删除成功",
//		})
//		return
//	}
//}

package router

import (
	"github.com/Pivot-Studio/Authorization-Template/pkg/ATutil"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreatTeam(context *gin.Context) {
	teamName:= context.PostForm("team_name")
	email,_:= ATutil.GetEmailFromCookie(context)
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
	teamID,_ := strconv.Atoi(context.PostForm("team_id"))
	is_team_exist := services.IsTeamExist(uint(teamID))
	if !is_team_exist {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "该组织不存在",
		})
		return
	}else{
		services.DeleteTeamInDB(uint(teamID))
		context.JSON(http.StatusOK, gin.H{
			"msg": "组织删除成功",
		})
		return
	}
}
func ChangeDetailOfTeam(context *gin.Context){
	teamID,_ := strconv.Atoi(context.PostForm("team_id"))
	newTeamName:=context.PostForm("new_team_name")
	mileStoneUpdate:=context.PostForm("mile_stone_update")
	descriptionUpdate:=context.PostForm("description_update")
	is_team_exist := services.IsTeamExist(uint(teamID))
	if !is_team_exist {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "该组织不存在",
		})
		return
	}else{
		services.ChangeTeamDetailInDB(uint(teamID),newTeamName,mileStoneUpdate,descriptionUpdate)
		context.JSON(http.StatusOK, gin.H{
			"msg": "修改组织信息成功",
		})
		return
	}

}
func GetTeamsList(context *gin.Context){
	startID, _ := strconv.Atoi(context.Query("start_id"))
	listSize, _ := strconv.Atoi(context.Query("list_size"))
	teams := services.GetTeams(uint(startID), listSize)
	results := []map[string]interface{}{}
	for _, value := range teams {
		userName:=services.GetUserNameByEmail(value.Owner_email)
		results = append(results, map[string]interface{}{
			"TeamName" :          value.Name,
			"Username":		  	  userName,
			"Description":        value.Description,
			"Mile_Stone":		  value.Mile_Stone,

		})

	}
	context.JSON(http.StatusOK, results)
	return
}
func GetTeam(context *gin.Context){
	teamId,_ := strconv.Atoi(context.Param("hole_id"))
	team,err:=services.GetTeamByID(uint(teamId))
	if err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "该组织不存在",
		})
		return
	}
	userName:=services.GetUserNameByEmail(team.Owner_email)
	result:=map[string]interface{}{
		"TeamName" :          team.Name,
		"Username":		  	  userName,
		"Description":        team.Description,
		"Mile_Stone":		  team.Mile_Stone,
	}
	context.JSON(http.StatusOK, result)
	return

}
package router

import (
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProjects(context *gin.Context){
	page,_:=strconv.Atoi(context.Query("page"))
	listSize,_:=strconv.Atoi(context.Query("enterprise_num"))
	projects:=services.GetProjects(page,listSize,true)
	results:=[]map[string]interface{}{}
	for _,temp:=range projects{
		results=append(results,map[string]interface{}{
			"projects_id":temp.ID,
			"projects_name":temp.Name,
			"projects_introduction":temp.BriefIntroduction,
		})
	}
	context.JSON(http.StatusOK,gin.H{
		"msg":results,
	})
}
func GetProjectDetails(context *gin.Context) {
	projectId, _ := strconv.Atoi(context.Param("project_id"))
	project, err := services.GetOneEnterprise(projectId)
	if err != nil {
		context.JSON(404, gin.H{`err`: err.Error()})
	}else{
		context.JSON(200,gin.H{
			"enterprise_id":project.ID,
			"enterprise_name":project.Name,
			"enterprise_introduction":project.BriefIntroduction,
			"enterprise_content":project.Content,
			"enterprise_email":project.Email,
			"enterprise_phone":project.Phone,
		})
	}
}
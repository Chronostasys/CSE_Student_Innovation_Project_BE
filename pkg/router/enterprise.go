package router

import (
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetEnterprises(context *gin.Context){
	page,_:=strconv.Atoi(context.Query("page"))
	listSize,_:=strconv.Atoi(context.Query("enterprise_num"))
	enterprises:=services.GetEnterprises(page,listSize,true)
	results:=[]map[string]interface{}{}
	for _,temp:=range enterprises{
		results=append(results,map[string]interface{}{
			"enterprise_id":temp.ID,
			"enterprise_name":temp.Name,
			"enterprise_introduction":temp.BriefIntroduction,
		})
	}
	context.JSON(http.StatusOK,gin.H{
		"msg":results,
	})
}


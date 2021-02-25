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

func GetEnterpriseDetails(context *gin.Context) {
	enterpriseId, _ := strconv.Atoi(context.Param("enterprise_id"))
	enterprise, err := services.GetOneEnterprise(enterpriseId)
	if err != nil {
		context.JSON(404, gin.H{`err`: err.Error()})
	}else{
		context.JSON(200,gin.H{
			"enterprise_id":enterprise.ID,
			"enterprise_name":enterprise.Name,
			"enterprise_introduction":enterprise.BriefIntroduction,
			"enterprise_content":enterprise.Content,
			"enterprise_email":enterprise.Email,
			"enterprise_phone":enterprise.Phone,
		})
	}


}


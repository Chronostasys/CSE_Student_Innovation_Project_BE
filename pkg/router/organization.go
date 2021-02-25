package router

import (
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/services"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateOrganization(context *gin.Context){
	otype:=context.PostForm("type")
	email:=context.PostForm("email")
	phone:=context.PostForm("phone")
	address:=context.PostForm("address")
	image:=context.PostForm("image")
	name:=context.PostForm("name")
	introduction:=context.PostForm("introduction")
	timestamp:=time.Now()
	useremail,_:=util.GetEmailFromToken(context)
	if otype=="enterprise"{
		info:=models.Enterprise{
			CreatTimestamp: timestamp,
			Address: address,
			Phone: phone,
			Email: email,
			Name: name,
			BriefIntroduction: introduction,
			CreaterEmail: useremail,
			Image: image,
		}
		err:=services.CreateEnterprise(info)
		if err!=nil{
			context.JSON(http.StatusInternalServerError,gin.H{
				"msg":"创建失败，请重试",
			})
			return
		}
		context.JSON(http.StatusOK,gin.H{
			"msg":"创建企业成功",
		})
		return
	}
	if otype=="project"{
		info:=models.Project{
			CreatTimestamp: timestamp,
			Address: address,
			Phone: phone,
			Email: email,
			Name: name,
			BriefIntroduction: introduction,
			CreaterEmail: useremail,
			Image: image,
		}
		err:=services.CreateProject(info)
		if err!=nil{
			context.JSON(http.StatusInternalServerError,gin.H{
				"msg":"创建失败，请重试",
			})
			return
		}
		context.JSON(http.StatusOK,gin.H{
			"msg":"创建企业成功",
		})
		return
	}
	context.JSON(http.StatusBadRequest,gin.H{
		"msg":"组织类型错误",
	})
	return
}
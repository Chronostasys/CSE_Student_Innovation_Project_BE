package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		user := api.Group("/auth")
		{
			user.Use(cors.Default())
			user.GET("/hello", Helloworld)
			user.POST("/signup", Signup)
			user.POST("/login", Login)
			user.POST("/sendVerifyCode", SendVerifyCode)
			user.POST("/changePassword_Email", ChangePassword_Email)
			user.POST("/changePasswordVerify_code", ChangePasswordVerify_code)
			user.POST("/changePassword", ChangePassword)

		}
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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
			user.POST("/changePassword_Email", ChangePasswordByEmail)
			user.POST("/changePasswordVerify_code", ChangePasswordVerifyCode)
			user.POST("/changePassword", ChangePassword)

		}
		blog := api.Group("/blog")
		{
			blog.Use(cors.Default())
			blog.GET("/getBlogNumber",GetBlogsNumber)
			blog.GET("",GetBlogs)
			blog.Use(Auth())
			blog.POST("", AddBlog)
			blog.DELETE("", DeleteBlog)
		}
		//team := api.Group("team")
		//{
		//	team.Use(cors.Default())
		//	team.GET("",GetTeamsList)
		//	team.GET("/:team_id",GetTeam)
		//	team.POST("", CreatTeam)
		//	team.DELETE("", DeleteTeam)
		//	team.PUT("", ChangeDetailOfTeam)
		//}
		//inviting := api.Group("inviting")
		//{
		//	inviting.GET("")
		//	inviting.POST("/invite",InviteTeamMember)
		//	inviting.POST("/accept",AcceptInvite)
		//}
		//myself:=api.Group("myself")
		//{
		//	myself.GET("/:email",GetMyTeam)
		//}
	}
}

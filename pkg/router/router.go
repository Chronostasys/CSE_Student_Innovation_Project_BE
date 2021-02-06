package router

import (
	"github.com/Pivot-Studio/Authorization-Template/pkg/ATrouter"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func InitRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		user := api.Group("/auth")
		{
			user.Use(cors.Default())
			user.POST("/isSessionExpired", ATrouter.IsTokenExpired)
			user.POST("/sendVerifyCode", ATrouter.SendVerifyCode)
			user.POST("/login", ATrouter.Login)
			user.POST("/changePassword", ATrouter.ForgetPassword)
			user.POST("/verifyCodeMatch", ATrouter.IsVerifyCodeMatch)
			user.GET("/hello", ATrouter.Helloworld)
			user.POST("/signup", ATrouter.Register)
		}
		blog := api.Group("/blog")
		{
			blog.POST("/add", AddBlog)
			blog.DELETE("/delete", DeleteBlog)
			blog.GET("/get",GetBlogs)
		}
		team := api.Group("team")
		{
			team.GET("",GetTeamsList)
			team.GET("/:team_id",GetTeam)
			team.POST("", CreatTeam)
			team.DELETE("", DeleteTeam)
			team.PUT("", ChangeDetailOfTeam)
		}
		inviting := api.Group("inviting")
		{
			inviting.GET("")
			inviting.POST("/invite",InviteTeamMember)
			inviting.POST("/accept",AcceptInvite)
		}
		myself:=api.Group("myself")
		{
			myself.GET("/:email",GetMyTeam)
		}
	}
}

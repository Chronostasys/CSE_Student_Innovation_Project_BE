package router

import (
	"github.com/Pivot-Studio/Authorization-Template/pkg/router"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth := api.Group("auth")
		{
			auth.GET("/hello", router.Helloworld)
		}
		blog := api.Group("/blog")
		{
			blog.POST("/add", AddBlog)
			blog.DELETE("/delete", DeleteBlog)
		}
		team := api.Group("team")
		{
			team.GET("",GetTeamsList)
			team.POST("", CreatTeam)
			team.DELETE("", DeleteTeam)
			team.PUT("", ChangeDetailOfTeam)
		}
		inviting := api.Group("inviting")
		{
			inviting.GET("")
			inviting.POST("", InviteTeamMember)
		}
	}
}
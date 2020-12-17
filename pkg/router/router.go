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
		blog := api.Group("blog")
		{
			blog.GET("")
		}
		team := api.Group("team")
		{
			team.GET("")
		}
		inviting := api.Group("inviting")
		{
			inviting.GET("")
			inviting.POST("",InviteTeamMember)
		}
	}

}

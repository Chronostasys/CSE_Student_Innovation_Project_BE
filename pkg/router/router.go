package router

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine){
	api:=r.Group("/api")
	{
		api.POST("/blog",AddBlog)
	}
}

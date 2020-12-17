package router

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine){
	api:=r.Group("/api")
	{
		blog:=api.Group("/blog")
		{
			blog.POST("/add",AddBlog)
			blog.DELETE("/delete",DeleteBlog)
		}
	}
}

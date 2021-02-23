package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func InitRouter(r *gin.Engine) {
	config:=cors.Config{
	    AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
	    AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type","access-control-allow-origin","token"},
	    AllowCredentials: false,
	}
	config.AllowAllOrigins=true
	r.Use(cors.New(config))
	api := r.Group("/api")
	{
		user := api.Group("/auth")
		{
			user.Use(cors.Default())
			user.GET("/hello", HelloWorld)
			user.POST("/signup", Signup)
			user.POST("/login", Login)
			user.POST("/sendVerifyCode", SendVerifyCode)
			user.POST("/changePasswordEmail", ChangePasswordByEmail)
			user.POST("/changePasswordVerifyCode", ChangePasswordVerifyCode)
			user.POST("/changePassword", ChangePassword)
			user.Use(Auth())
			user.GET("/myself",GetMyselfInfo)
		}
		blog := api.Group("/blog")
		{
			blog.Use(cors.Default())
			blog.GET("/getBlogNumber",GetBlogsNumber)
			blog.GET("",GetBlogs)
			blog.GET("/detail/:blog_id",GetBlog)
			blog.Use(Auth())
			blog.Use(cors.Default())
			blog.POST("", AddBlog)
			blog.DELETE("", DeleteBlog)
		}
		enterprise := api.Group("/enterprise")
		{
			enterprise.GET("/enterpriseList", GetEnterprises)
			enterprise.GET("/enterpriseDetail", )
		}
		project := api.Group("/project")
		{
			project.GET("/projectList", )
			project.GET("/projectDetail", )
		}
	}
}

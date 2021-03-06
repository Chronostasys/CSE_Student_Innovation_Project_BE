package main

import (
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8080")
}

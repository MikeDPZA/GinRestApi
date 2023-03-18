package main

import (
	"GinRestApi/pkg/controllers"
	"GinRestApi/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controllers.AddApiController(r)

	db.Init()

	panic(r.Run(":5001"))
}

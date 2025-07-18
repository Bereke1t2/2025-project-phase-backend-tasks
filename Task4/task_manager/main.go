package main

import (
	"github.com/gin-gonic/gin"
	"task_manager/router"
)

func main() {
	route := gin.Default()

	router.SetUp_Router(route)

	if err := route.Run(":8000"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
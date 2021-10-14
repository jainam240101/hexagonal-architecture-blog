package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jainam240101/todo/db"
	Routes "github.com/jainam240101/todo/server/routes"
)

func main() {
	db.Init()
	router := gin.Default()
	Routes.UserRoutes(&router.RouterGroup, db.DB)
	router.Run()
}

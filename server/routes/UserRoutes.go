package routes

import (
	"github.com/gin-gonic/gin"
	domain "github.com/jainam240101/todo/server/domain/Users"
	service "github.com/jainam240101/todo/server/services"
	handlers "github.com/jainam240101/todo/server/handlers"
	"gorm.io/gorm"
)

func UserRoutes(apiRouter *gin.RouterGroup, db *gorm.DB) {
	uh := handlers.Userhandlers{Service: service.NewCustomerService(domain.NewUserRepositoryDb(db))}
	route := apiRouter.Group("/users")
	{
		route.POST("/", uh.CreateUser)
	}
}

package users

import (
	"github.com/gin-gonic/gin"
	domain "github.com/jainam240101/todo/server/domain/Users"
	service "github.com/jainam240101/todo/server/services"
)

type Userhandlers struct {
	Service service.UserService
}

func (uh *Userhandlers) CreateUser(c *gin.Context) {
	u := domain.UserModel{}
	if err := c.BindJSON(&u); err != nil {
		c.JSON(406, gin.H{
			"message": "Body Not Proper",
		})
		return
	}
	user, err := uh.Service.CreateUser(u)
	if err != nil {
		c.JSON(406, gin.H{
			"message": "Body Not Proper",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": user,
	})
}

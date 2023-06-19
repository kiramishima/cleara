package users

import (
	"cleara/internal/core/ports"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	userService ports.UserService
}

func NewUserHandlers(userService ports.UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (userSrv *UserHandlers) GetProfile(c *gin.Context) {
	userProfile, err := userSrv.userService.GetProfile(c.Param("id"))
	if err != nil {
		// throw error
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(userProfile)
	c.JSON(200, userProfile)
}

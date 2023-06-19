package users

import (
	"cleara/internal/core/ports"
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
	println(userProfile)
	c.JSON(200, userProfile)
}

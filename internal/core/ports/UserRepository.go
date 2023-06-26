package ports

import (
	"cleara/internal/core/domain"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	GetProfile(id string) (*domain.User, error)
}

type UserService interface {
	GetProfile(id string) (*domain.User, error)
}

type UserHandlers interface {
	GetProfile(c *gin.Context)
}

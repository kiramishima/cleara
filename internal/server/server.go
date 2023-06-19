package server

import (
	"cleara/internal/core/ports"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	userHandlers ports.UserHandlers
}

func NewServer(uHandlers ports.UserHandlers) *Server {
	return &Server{
		userHandlers: uHandlers,
	}
}

func (s *Server) Initialize() {
	app := gin.New()
	v1 := app.Group("/v1")

	userRoutes := v1.Group("/users")
	userRoutes.GET("/:id", s.userHandlers.GetProfile)

	err := app.Run(":5000")
	if err != nil {
		log.Fatal(err)
	}
}

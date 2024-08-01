package server

import (
	"hackerNewsApi/internal/components/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin *gin.Engine
}

func NewGin(config *config.Config) *Server {
	var gin = gin.Default()
	return &Server{Gin: gin}
}

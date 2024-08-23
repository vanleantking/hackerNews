package server

import (
	"fmt"
	"hackerNewsApi/pkg/config"

	"github.com/gin-gonic/gin"
)

type serverInstance struct {
	Gin    *gin.Engine
	Config *config.Config
}

type Server interface {
	Run()
	GetConfig() *config.Config
	GetEngine() *gin.Engine
}

func NewServer(config *config.Config) Server {
	var gin = gin.Default()
	return &serverInstance{Gin: gin, Config: config}
}

func (s *serverInstance) Run() {
	address := fmt.Sprintf("%s:%s", s.Config.ServerHost, s.Config.APIPort)
	fmt.Println("address, ", address)
	s.Gin.Run(address)
}

func (s *serverInstance) GetEngine() *gin.Engine {
	return s.Gin
}

func (s *serverInstance) GetConfig() *config.Config {
	return s.Config
}

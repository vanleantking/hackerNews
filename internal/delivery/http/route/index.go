package route

import (
	"net/http"
	"time"

	"hackerNewsApi/internal/components/config"
	"hackerNewsApi/internal/components/server"

	"github.com/gin-gonic/gin"
)

func Setup(env *config.Config, timeout time.Duration, srv *server.Server) {
	publicRouter := srv.Gin.Group("")
	// All Public APIs
	NewIndexRouter(env, timeout, publicRouter)
}

func NewIndexRouter(env *config.Config, timeout time.Duration, group *gin.RouterGroup) {
	group.GET("/", Index)
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, "ping")

}

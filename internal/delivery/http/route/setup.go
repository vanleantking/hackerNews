package route

import (
	"hackerNewsApi/pkg/gorm"
	"hackerNewsApi/pkg/logger"
	"hackerNewsApi/pkg/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	Server     server.Server
	Logger     logger.Logger
	APIVersion *gin.RouterGroup
	DB         gorm.Database
}

func NewRouteConfig(srv server.Server, logger logger.Logger, db gorm.Database) *RouteConfig {
	return &RouteConfig{
		Server: srv,
		Logger: logger,
		DB:     db,
	}
}

func (routeConfig *RouteConfig) Setup() {
	// setup version api
	apiversion := routeConfig.Server.GetConfig().APIVersion
	routeConfig.APIVersion = routeConfig.Server.GetEngine().Group("/" + apiversion)
	// All Public APIs
	routeConfig.SetupGuestRoute()

	routeConfig.HNAPIRouter()
}

func (routeConfig *RouteConfig) SetupGuestRoute() {
	publicRouter := routeConfig.APIVersion.Group("")
	NewIndexRouter(publicRouter)
}

func NewIndexRouter(group *gin.RouterGroup) {
	group.GET("/", Index)
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, "ping")

}

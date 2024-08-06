package route

import (
	"hackerNewsApi/internal/components/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App        server.Server
	APIVersion *gin.RouterGroup
}

func NewRouteConfig(app server.Server) *RouteConfig {
	return &RouteConfig{
		App: app,
	}
}

func (routeConfig *RouteConfig) Setup() {
	// setup version api
	apiversion := routeConfig.App.GetConfig().APIVersion
	routeConfig.APIVersion = routeConfig.App.GetEngine().Group("/" + apiversion)
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

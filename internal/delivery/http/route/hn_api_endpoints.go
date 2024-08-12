package route

import (
	"hackerNewsApi/internal/delivery/http/controller/frontend"
	service "hackerNewsApi/internal/service/hn_api"
	"hackerNewsApi/pkg/request"
)

func (routeConfig *RouteConfig) HNAPIRouter() {
	hnApiService := service.NewHNAPIClient(
		request.NewClient(),
		routeConfig.App.GetConfig().HNBaseURL,
		routeConfig.App.GetConfig().HNAPIVersion,
		routeConfig.App.GetConfig().HNAPIFormat,
	)
	hnListItemController := frontend.NewListTopStoriesController(hnApiService)
	hnAPIRouter := routeConfig.APIVersion.Group("hn-api")
	hnAPIRouter.POST("/top-stories", hnListItemController.ListTopStories)
}

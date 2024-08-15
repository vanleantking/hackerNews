package route

import (
	"hackerNewsApi/internal/delivery/http/controller/frontend"
	"hackerNewsApi/internal/repository"
	service "hackerNewsApi/internal/service/hn_api"
	"hackerNewsApi/internal/usecase"
	"hackerNewsApi/pkg/request"
)

func (routeConfig *RouteConfig) HNAPIRouter() {
	hnApiService := service.NewHNAPIClient(
		request.NewClient(),
		routeConfig.Server.GetConfig().HNBaseURL,
		routeConfig.Server.GetConfig().HNAPIVersion,
		routeConfig.Server.GetConfig().HNAPIFormat,
	)

	// get top-stories hacker news endpoint api
	itemRepo := repository.NewListItemRepository(&routeConfig.Logger, routeConfig.DB.GetDb())
	listItemUsc := usecase.NewListItemUsercase(itemRepo)
	hnListItemController := frontend.NewListTopStoriesController(hnApiService, listItemUsc)
	hnAPIRouter := routeConfig.APIVersion.Group("hn-api")
	hnAPIRouter.POST("/top-stories", hnListItemController.ListTopStories)

	// get item detail hacker news endpoint api
	itemDetailRepo := repository.NewItemDetailRepository(&routeConfig.Logger, routeConfig.DB.GetDb())
	detailItemUsc := usecase.NewItemDetailUseCase(itemDetailRepo)
	itemDetailController := frontend.NewItemDetailController(hnApiService, detailItemUsc)
	hnAPIRouter.POST("/item-detail", itemDetailController.GetDetailItem)
}

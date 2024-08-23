package route

import (
	"hackerNewsApi/internal/delivery/http/controller/frontend"
	"hackerNewsApi/internal/infrastructure/repository/postgre"
	service "hackerNewsApi/internal/infrastructure/service/hn_api"
	"hackerNewsApi/internal/infrastructure/usecase"
	"hackerNewsApi/pkg/request"
)

const (
	HNAPIGroup = "hn-api"
)

func (routeConfig *RouteConfig) HNAPIRouter() {
	hnApiService := service.NewHNAPIClient(
		request.NewClient(),
		routeConfig.Server.GetConfig().HNBaseURL,
		routeConfig.Server.GetConfig().HNAPIVersion,
		routeConfig.Server.GetConfig().HNAPIFormat,
	)

	// get top-stories hacker news endpoint api
	itemRepo := postgre.NewListItemRepository(&routeConfig.Logger, routeConfig.DB.GetDb())
	listItemUsc := usecase.NewListItemUsercase(itemRepo)
	hnListItemController := frontend.NewListTopStoriesController(hnApiService, listItemUsc)
	hnAPIRouter := routeConfig.APIVersion.Group(HNAPIGroup)
	hnAPIRouter.POST("/top-stories", hnListItemController.ListTopStories)

	// get item detail hacker news endpoint api
	itemDetailRepo := postgre.NewItemDetailRepository(&routeConfig.Logger, routeConfig.DB.GetDb())
	detailItemUsc := usecase.NewItemDetailUseCase(itemDetailRepo)
	itemDetailController := frontend.NewItemDetailController(hnApiService, detailItemUsc)
	hnAPIRouter.POST("/item-detail", itemDetailController.GetDetailItem)
}

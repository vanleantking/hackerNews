package frontend

import (
	"hackerNewsApi/internal/delivery/http/httpstatus"
	"hackerNewsApi/internal/delivery/http/params"
	"hackerNewsApi/internal/domains/services"
	"hackerNewsApi/internal/domains/usecases"
	"hackerNewsApi/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HNItemsController struct {
	APIHNService services.HNAPIClient
	ListItemUsc  usecases.ListItemUseCase
}

func NewListTopStoriesController(
	apiService services.HNAPIClient,
	listItemUsc usecases.ListItemUseCase,
) *HNItemsController {
	return &HNItemsController{
		APIHNService: apiService,
		ListItemUsc:  listItemUsc,
	}
}

func (hnItemController *HNItemsController) ListTopStories(c *gin.Context) {
	var listItemParams params.ListItemParams

	err := c.BindJSON(&listItemParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpstatus.ErrorResponse{Message: err.Error()})
		return
	}
	errs := params.ValidatorListItemsRequest(listItemParams)
	if errs != nil {
		err = errs[0]
		c.JSON(http.StatusBadRequest, httpstatus.ErrorResponse{Message: err.Error()})
		return
	}

	items, err := hnItemController.APIHNService.GetListTopStories(
		listItemParams.Method,
		listItemParams.Params,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpstatus.ErrorResponse{Message: err.Error()})
		return
	}
	entities := model.MapperItemsCreateEntity(items.Items)
	err = hnItemController.ListItemUsc.InsertBulkTopStoriesV2(entities)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpstatus.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, httpstatus.Response[interface{}]{Message: httpstatus.Success})
}

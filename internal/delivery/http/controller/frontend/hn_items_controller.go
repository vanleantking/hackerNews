package frontend

import (
	"fmt"
	"hackerNewsApi/internal/delivery/http/errors"
	"hackerNewsApi/internal/delivery/http/params"
	hnAPIService "hackerNewsApi/internal/service/hn_api"
	"hackerNewsApi/internal/service/hn_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HNItemsController struct {
	APIHNService hnAPIService.HNAPIClient
}

func NewListTopStoriesController(apiService hnAPIService.HNAPIClient) *HNItemsController {
	return &HNItemsController{
		APIHNService: apiService,
	}
}

func (hnItemController *HNItemsController) ListTopStories(c *gin.Context) {
	var listItemParams params.ListItemParams

	err := c.BindJSON(&listItemParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{Message: err.Error()})
		return
	}
	errs := params.ValidatorListItemsRequest(listItemParams)
	if errs != nil {
		err = errs[0]
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{Message: err.Error()})
		return
	}

	items, err := hnItemController.APIHNService.GetListTopStories(listItemParams.Method, listItemParams.Params)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{Message: err.Error()})
		return
	}
	entities := model.MapperItemsCreateEntity(items.Items)
	fmt.Println("entities, ", entities, items, len(*entities), len(items.Items))
}

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

type HNItemDetailController struct {
	APIHNService  services.HNAPIClient
	ItemDetailUsc usecases.ItemDetailUseCase
}

func NewItemDetailController(
	apiService services.HNAPIClient,
	itemDetailUsc usecases.ItemDetailUseCase,
) *HNItemDetailController {
	return &HNItemDetailController{
		APIHNService:  apiService,
		ItemDetailUsc: itemDetailUsc,
	}
}

func (hnItemController *HNItemDetailController) GetDetailItem(c *gin.Context) {
	var itemDetailParams params.ItemDetailParams

	err := c.BindJSON(&itemDetailParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpstatus.ErrorResponse{Message: err.Error()})
		return
	}
	errs := params.ValidatorItemDetailRequest(itemDetailParams)
	if errs != nil {
		err = errs[0]
		c.JSON(http.StatusBadRequest, httpstatus.ErrorResponse{Message: err.Error()})
		return
	}

	item, err := hnItemController.APIHNService.GetItemDetailById(
		itemDetailParams.Method,
		itemDetailParams.ItemID,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpstatus.ErrorResponse{Message: err.Error()})
		return
	}
	itemEntity := model.MapperSingleItemEntity(*item)
	err = hnItemController.ItemDetailUsc.UpdateDetailTopStory(itemEntity)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpstatus.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, httpstatus.Response[interface{}]{Message: httpstatus.Success})
}

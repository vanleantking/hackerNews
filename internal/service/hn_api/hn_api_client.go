package service

import (
	"encoding/json"
	"fmt"
	"hackerNewsApi/internal/service/hn_api/common"
	"hackerNewsApi/internal/service/hn_api/model"
	"hackerNewsApi/pkg/request"
)

type hnAPIClient struct {
	ApiClient  *request.Client
	BaseURL    string
	VersionAPI string
	APIFormat  string
}

type HNAPIClient interface {
	GetListTopStories(method string, params map[string]interface{}) (*model.HNItems, error)
	GetItemDetailById(method string, itemId int) (*model.Item, error)
}

func NewHNAPIClient(
	client *request.Client,
	apiBaseURL, apiVersion, apiFormat string) HNAPIClient {
	return &hnAPIClient{
		ApiClient:  client,
		BaseURL:    apiBaseURL,
		VersionAPI: apiVersion,
		APIFormat:  apiFormat,
	}
}

func (api *hnAPIClient) GetListTopStories(method string, params map[string]interface{}) (*model.HNItems, error) {
	endPoint := common.ENDPOINT_TOPSTORIES
	urlRequest := api.generateFullURLRequest(endPoint, 0)
	resByte, err := api.ApiClient.MakeRequest(method, urlRequest, params)
	if err != nil {
		return new(model.HNItems), err
	}
	var items []int64
	err = json.Unmarshal(resByte, &items)
	var tmp = make([]int64, len(items))
	copy(tmp, items)
	return &model.HNItems{Items: tmp}, nil
}

func (api *hnAPIClient) GetItemDetailById(method string, itemId int) (*model.Item, error) {
	endPoint := common.ENDPOINT_ITEM
	urlRequest := api.generateFullURLRequest(endPoint, itemId)
	fmt.Println("urlRequest GetItemDetailById, ", urlRequest, method, itemId, endPoint)
	var params = map[string]interface{}{
		"item_id": itemId,
	}
	resByte, err := api.ApiClient.MakeRequest(method, urlRequest, params)
	if err != nil {
		return new(model.Item), err
	}
	var item model.Item
	err = json.Unmarshal(resByte, &item)
	return &item, err
}

func (api *hnAPIClient) generateFullURLRequest(endPoint string, itemId int) string {
	// on end-point get item detail
	if endPoint == common.ENDPOINT_ITEM {
		return fmt.Sprintf(
			"%s/%s/%s/%d.%s",
			api.BaseURL,
			api.VersionAPI,
			endPoint,
			itemId,
			api.APIFormat,
		)
	}
	// otherwise other endpoint
	return fmt.Sprintf(
		"%s/%s/%s.%s",
		api.BaseURL,
		api.VersionAPI,
		endPoint,
		api.APIFormat,
	)
}

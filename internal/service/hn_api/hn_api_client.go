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
	GetListTopStories(method string, params map[string]string) (*model.HNItems, error)
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

func (api *hnAPIClient) GetListTopStories(method string, params map[string]string) (*model.HNItems, error) {
	endPoint := common.ENDPOINT_TOPSTORIES
	urlRequest := api.generateFullURLRequest(endPoint)
	fmt.Println("urlRequest GetListTopStories, ", urlRequest)
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

func (api *hnAPIClient) generateFullURLRequest(endPoint string) string {
	return fmt.Sprintf(
		"%s/%s/%s.%s",
		api.BaseURL,
		api.VersionAPI,
		endPoint,
		api.APIFormat,
	)
}
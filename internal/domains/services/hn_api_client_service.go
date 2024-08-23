package services

import "hackerNewsApi/internal/model"

type HNAPIClient interface {
	GetListTopStories(method string, params map[string]interface{}) (*model.HNItems, error)
	GetItemDetailById(method string, itemId int) (*model.Item, error)
}

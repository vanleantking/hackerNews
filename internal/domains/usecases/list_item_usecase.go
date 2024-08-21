package usecases

import (
	"hackerNewsApi/internal/domains/entity"
	"hackerNewsApi/internal/model"
)

type ListItemRepository interface {
	UpsertListItems(items []entity.Item) []error
	UpsertBulkItems(items []entity.Item) error
	FindItemListUpdate(conditions map[string]interface{}) ([]entity.Item, error)
}

type ListItemUseCase interface {
	InsertBulkTopStories(items []entity.Item) []error
	InsertBulkTopStoriesV2(items []entity.Item) error
	FindItemsUpdate(conditions map[string]interface{}) ([]model.Item, error)
}

package usecases

import "hackerNewsApi/internal/domains/entity"

type ListItemRepository interface {
	UpsertListItems(items []entity.Item) []error
	UpsertBulkItems(items []entity.Item) error
}

type ListItemUseCase interface {
	InsertBulkTopStories(items []entity.Item) []error
	InsertBulkTopStoriesV2(items []entity.Item) error
}

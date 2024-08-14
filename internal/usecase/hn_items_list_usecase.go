package usecase

import (
	"hackerNewsApi/internal/entity"
	"hackerNewsApi/internal/repository"
)

type listItemsUsecase struct {
	ListItemRepo repository.ListItemRepository
}

type ListItemUseCase interface {
	InsertBulkTopStories(items []entity.Item) []error
	InsertBulkTopStoriesV2(items []entity.Item) error
}

func NewListItemUsercase(itemsRepo repository.ListItemRepository) ListItemUseCase {
	return &listItemsUsecase{
		ListItemRepo: itemsRepo,
	}
}

func (itemsUsc *listItemsUsecase) InsertBulkTopStories(items []entity.Item) []error {
	// mapper items into entities
	return itemsUsc.ListItemRepo.UpsertListItems(items)
}

func (itemsUsc *listItemsUsecase) InsertBulkTopStoriesV2(items []entity.Item) error {
	// mapper items into entities
	return itemsUsc.ListItemRepo.UpsertBulkItems(items)
}

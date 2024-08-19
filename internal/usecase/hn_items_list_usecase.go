package usecase

import (
	"hackerNewsApi/internal/domains/entity"
	"hackerNewsApi/internal/domains/usecases"
)

type listItemsUsecase struct {
	ListItemRepo usecases.ListItemRepository
}

func NewListItemUsercase(itemsRepo usecases.ListItemRepository) usecases.ListItemUseCase {
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

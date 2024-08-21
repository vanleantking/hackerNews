package usecase

import (
	"hackerNewsApi/internal/domains/entity"
	"hackerNewsApi/internal/domains/usecases"
	"hackerNewsApi/internal/infrastructure/mapper"
	"hackerNewsApi/internal/model"
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

func (itemUsc *listItemsUsecase) FindItemsUpdate(conditions map[string]interface{}) ([]model.Item, error) {
	var items = make([]model.Item, 0)
	entities, er := itemUsc.ListItemRepo.FindItemListUpdate(conditions)
	if er != nil {
		return items, er
	}
	// mapper entities into items
	for _, entityItem := range entities {
		item := mapper.MapperEntityItem(entityItem)
		items = append(items, item)
	}
	var tmp = make([]model.Item, len(entities))
	copy(tmp, items)
	return tmp, nil
}

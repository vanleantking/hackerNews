package usecase

import (
	"hackerNewsApi/internal/domains/entity"
	"hackerNewsApi/internal/domains/usecases"
)

type itemDetailUsecase struct {
	ItemDetailRepository usecases.ItemDetailRepository
}

func NewItemDetailUseCase(itemDetailRepo usecases.ItemDetailRepository) usecases.ItemDetailUseCase {
	return &itemDetailUsecase{
		ItemDetailRepository: itemDetailRepo,
	}
}

func (itemUsc *itemDetailUsecase) UpdateDetailTopStory(item entity.Item) error {
	// mapper items into entities
	return itemUsc.ItemDetailRepository.UpdateItem(item)
}

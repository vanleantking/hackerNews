package usecase

import (
	"hackerNewsApi/internal/entity"
	"hackerNewsApi/internal/repository"
)

type itemDetailUsecase struct {
	ItemDetailRepository repository.ItemDetailRepository
}

type ItemDetailUseCase interface {
	UpdateDetailTopStory(items entity.Item) error
}

func NewItemDetailUseCase(itemDetailRepo repository.ItemDetailRepository) ItemDetailUseCase {
	return &itemDetailUsecase{
		ItemDetailRepository: itemDetailRepo,
	}
}

func (itemUsc *itemDetailUsecase) UpdateDetailTopStory(item entity.Item) error {
	// mapper items into entities
	return itemUsc.ItemDetailRepository.UpdateItem(item)
}

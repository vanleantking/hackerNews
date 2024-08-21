package usecases

import (
	"hackerNewsApi/internal/domains/entity"
)

type ItemDetailRepository interface {
	UpdateItem(items entity.Item) error
}

type ItemDetailUseCase interface {
	UpdateDetailTopStory(item entity.Item) error
}

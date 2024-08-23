package mapper

import (
	"hackerNewsApi/internal/domains/entity"
	"hackerNewsApi/internal/model"
)

func MapperEntityItem(entityItem entity.Item) model.Item {
	return model.Item{
		ItemID:  entityItem.HNItemID,
		ItemURL: entityItem.URL,
	}
}

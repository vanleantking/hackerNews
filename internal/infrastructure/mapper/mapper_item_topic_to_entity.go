package mapper

import (
	"hackerNewsApi/internal/domains/entity"
	"hackerNewsApi/internal/infrastructure/pubsub/item"
)

func MapperItemTopicEntity(itemTopic item.PubsubHNItemDetail) entity.Item {
	return entity.Item{
		Score:       itemTopic.Score,
		By:          itemTopic.By,
		Title:       itemTopic.ItemType,
		URL:         itemTopic.URL,
		DescenDants: itemTopic.DescenDants,
		Kids:        itemTopic.Kids,
		ItemType:    itemTopic.ItemType,
		ItemDeleted: itemTopic.ItemDeleted,
		ItemStatus:  itemTopic.ItemStatus,
		UpdatedAt:   itemTopic.UpdatedAt,
	}
}

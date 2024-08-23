package mapper

import (
	"hackerNewsApi/internal/domains/entity"
	"hackerNewsApi/internal/model/item"
)

func MapperItemTopicEntity(itemTopic item.PubsubHNItemDetail) entity.Item {
	return entity.Item{
		HNItemID:    itemTopic.HNID,
		Text:        itemTopic.ItemText,
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

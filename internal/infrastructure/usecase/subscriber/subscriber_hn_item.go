package subscriber

import (
	"encoding/json"
	"errors"
	"fmt"
	"hackerNewsApi/internal/domains/usecases"
	"hackerNewsApi/internal/domains/usecases/subscribers"
	"hackerNewsApi/internal/infrastructure/mapper"
	"hackerNewsApi/internal/model/item"
)

type subscriberHNItemUsc struct {
	ItemDetailRepository usecases.ItemDetailRepository
}

func NewSubscriberHNItemUsc(itemDetail usecases.ItemDetailRepository) subscribers.SubscriberItemDetailUsc {
	return &subscriberHNItemUsc{
		ItemDetailRepository: itemDetail,
	}
}

func (sub *subscriberHNItemUsc) SubscriberUpdateDetailTopStory() func(data []byte) error {
	return func(data []byte) error {
		fmt.Println("--------------SubscriberUpdateDetailTopStory, ,", string(data))
		var itemDetail item.PubsubHNItemDetail
		err := json.Unmarshal(data, &itemDetail)
		// map from topic into entity for update the item
		if err != nil {
			return errors.New("Error, can not assert data PubSubEvent into PubsubHNItemDetail item")
		}
		entityItem := mapper.MapperItemTopicEntity(itemDetail)
		err = sub.ItemDetailRepository.UpdateItem(entityItem)
		return err
	}
}

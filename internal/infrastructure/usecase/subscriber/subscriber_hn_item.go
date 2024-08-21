package subscriber

import (
	"errors"
	"hackerNewsApi/internal/domains/usecases"
	"hackerNewsApi/internal/domains/usecases/subscriber"
	"hackerNewsApi/internal/infrastructure/mapper"
	"hackerNewsApi/internal/infrastructure/pubsub"
	"hackerNewsApi/internal/infrastructure/pubsub/item"
)

type hnItemProfileSub struct {
	ItemDetailRepository usecases.ItemDetailRepository
}

func NewTiktokProfileSub(itemDetail usecases.ItemDetailRepository) subscriber.SubscriberItemDetailUseCase {
	return &hnItemProfileSub{
		ItemDetailRepository: itemDetail,
	}
}

func (sub *hnItemProfileSub) SubscriberUpdateDetailTopStory() func(data pubsub.PubSubEvent) error {
	return func(data pubsub.PubSubEvent) error {
		// map from topic into entity for update the item
		itemDetail, ok := data.(item.PubsubHNItemDetail)
		if !ok {
			return errors.New("Error, can not assert data PubSubEvent into PubsubHNItemDetail item")
		}
		entityItem := mapper.MapperItemTopicEntity(itemDetail)

		if err := sub.ItemDetailRepository.UpdateItem(entityItem); err != nil {
			return err
		}
		return nil
	}
}

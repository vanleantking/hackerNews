package batchprocessing

import (
	"context"
	"fmt"
	"hackerNewsApi/internal/common"
	"hackerNewsApi/internal/domains/pubsub"
	"hackerNewsApi/internal/domains/services"
	"hackerNewsApi/internal/domains/usecases"
	hnCommon "hackerNewsApi/internal/infrastructure/service/hn_api/common"
	"hackerNewsApi/internal/model"
	itemModel "hackerNewsApi/internal/model/item"
)

type BatchProcessingItemList interface {
	ProcessItemList() error
	ProcessUpdateHNItems() error
	RunBatchProcess() error
}

type batchProcessingTopStories struct {
	ItemListUsc  usecases.ListItemUseCase
	APIHNService services.HNAPIClient
	PubSub       pubsub.RedisPublish
}

func NewBatchProcessTopStories(
	itemListUsc usecases.ListItemUseCase,
	apiHNService services.HNAPIClient,
	pubRedis pubsub.RedisPublish,
) BatchProcessingItemList {
	return &batchProcessingTopStories{
		ItemListUsc:  itemListUsc,
		APIHNService: apiHNService,
		PubSub:       pubRedis,
	}
}

func (batchProcess *batchProcessingTopStories) ProcessItemList() error {
	var paramsTopStories = map[string]interface{}{
		"print": "pretty",
	}
	items, err := batchProcess.APIHNService.GetListTopStories(
		common.HTTPGet,
		paramsTopStories,
	)
	if err != nil {
		return err
	}
	entities := model.MapperItemsCreateEntity(items.Items)
	return batchProcess.ItemListUsc.InsertBulkTopStoriesV2(entities)
}

func (batchProcess *batchProcessingTopStories) ProcessUpdateHNItems() error {
	var paramsItemsUpdate = map[string]interface{}{
		"item_status": hnCommon.ITEM_STATUS_NEW,
	}
	itemsHN, err := batchProcess.ItemListUsc.FindItemsUpdate(paramsItemsUpdate)
	fmt.Println("itemsHN, err, ", itemsHN, err)
	if err != nil {
		return err
	}
	// for each item, get the detail from hn-api client
	for _, item := range itemsHN {
		err := batchProcess.publisherSingleItemUpdate(item.ItemID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (batchProcess *batchProcessingTopStories) RunBatchProcess() error {
	fmt.Println("-------- start RunBatchProcess--------")
	err := batchProcess.ProcessItemList()
	if err != nil {
		return err
	}
	return batchProcess.ProcessUpdateHNItems()
}

func (batchProcess *batchProcessingTopStories) publisherSingleItemUpdate(itemId uint) error {
	hnItem, err := batchProcess.APIHNService.GetItemDetailById(common.HTTPGet, int(itemId))
	if err != nil {
		return err
	}
	itemTopic := itemModel.MapperItemToPubsubItem(hnItem)
	byteRepresenter, err := itemTopic.BytePresenter()
	if err != nil {
		return err
	}
	return batchProcess.PubSub.Publish(
		context.Background(),
		itemTopic.GetTopicName(),
		byteRepresenter,
	)
}

package batchprocessing

import (
	"hackerNewsApi/internal/common"
	"hackerNewsApi/internal/domains/services"
	"hackerNewsApi/internal/domains/usecases"
	hnCommon "hackerNewsApi/internal/infrastructure/service/hn_api/common"
	"hackerNewsApi/internal/model"
)

type BatchProcessingItemList interface {
	ProcessItemList() error
	ProcessUpdateHNItems() error
}

type batchProcessingTopStories struct {
	ItemListUsc   usecases.ListItemUseCase
	ItemDetailUsc usecases.ItemDetailUseCase
	APIHNService  services.HNAPIClient
}

func NewBatchProcessTopStories(
	itemListUsc usecases.ListItemUseCase,
	itemDetailUsc usecases.ItemDetailUseCase,
	apiHNService services.HNAPIClient,
) BatchProcessingItemList {
	return &batchProcessingTopStories{
		ItemListUsc:   itemListUsc,
		ItemDetailUsc: itemDetailUsc,
		APIHNService:  apiHNService,
	}
}

func (batchProcess *batchProcessingTopStories) ProcessItemList() error {
	err := batchProcess.processUpdateListStories()
	if err != nil {
		return err
	}
	return err
}

// processUpdateListStories: up-to-date list of newest top-stories from api-hn first
func (batchProcess *batchProcessingTopStories) processUpdateListStories() error {
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
	err = batchProcess.ItemListUsc.InsertBulkTopStoriesV2(entities)
	return err
}

func (batchProcess *batchProcessingTopStories) ProcessUpdateHNItems() error {
	var paramsItemsUpdate = map[string]interface{}{
		"status": hnCommon.ITEM_STATUS_NEW,
	}
	itemsHN, err := batchProcess.ItemListUsc.FindItemsUpdate(paramsItemsUpdate)
	if err != nil {
		return err
	}
	// for each item, get the detail from hn-api client
	for _, item := range itemsHN {
		err := batchProcess.processSingleItemUpdate(item.ItemID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (batchProcess *batchProcessingTopStories) processSingleItemUpdate(itemId uint) error {
	item, err := batchProcess.APIHNService.GetItemDetailById(common.HTTPGet, int(itemId))
	if err != nil {
		return err
	}
	itemEntity := model.MapperSingleItemEntity(*item)
	return batchProcess.ItemDetailUsc.UpdateDetailTopStory(itemEntity)
}

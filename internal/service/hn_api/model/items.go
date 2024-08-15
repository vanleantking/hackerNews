package model

import (
	"encoding/json"
	"hackerNewsApi/internal/entity"
	"hackerNewsApi/internal/service/hn_api/common"
	"time"
)

type Item struct {
	ItemBy      string          `json:"by"`
	ItemID      uint            `json:"id"`
	ItemScore   int             `json:"score"`
	ItemText    string          `json:"text"`
	ItemTime    uint            `json:time"`
	ItemTitle   string          `json:"title"`
	ItemType    string          `json:"type"`
	ItemURL     string          `json:"url"`
	DescenDants uint            `json:"descendants"`
	Kids        json.RawMessage `json:"kids"`
}

type HNItems struct {
	Items []int64
}

func MapperItemsCreateEntity(items []int64) []entity.Item {
	currentTime := time.Now()
	var tmp = make([]entity.Item, 0)
	for _, item := range items {
		itemEntity := entity.Item{
			HNItemID:  uint(item),
			Score:     0,
			UpdatedAt: currentTime.Unix(),
			CreatedAt: currentTime.Unix(),
		}
		tmp = append(tmp, itemEntity)
	}
	if len(tmp) > 0 {
		result := make([]entity.Item, len(tmp))
		copy(result, tmp)
		return result
	}
	return nil
}

func MapperItemsUpsertEntity(items []Item) *[]entity.Item {
	currentTime := time.Now()
	var tmp = make([]entity.Item, 0)
	for _, item := range items {
		itemEntity := entity.Item{}
		itemEntity.HNItemID = uint(item.ItemID)
		itemEntity.URL = item.ItemURL
		itemEntity.Text = item.ItemText
		itemEntity.Title = item.ItemTitle
		itemEntity.By = item.ItemBy
		itemEntity.ItemType = item.ItemType
		itemEntity.Category = common.ITEM_CATEGORY_DEFAULT
		itemEntity.Score = int(item.ItemScore)
		itemEntity.CreatedTime = int64(item.ItemTime)
		itemEntity.UpdatedAt = currentTime.Unix()
		itemEntity.CreatedAt = currentTime.Unix()
		itemEntity.ItemStatus = common.ITEM_STATUS_NEW
		tmp = append(tmp, itemEntity)
	}
	if len(tmp) > 0 {
		result := make([]entity.Item, len(tmp))
		copy(result, tmp)
		return &result
	}
	return nil
}

func MapperSingleItemEntity(item Item) entity.Item {
	currentTime := time.Now()
	itemEntity := entity.Item{
		HNItemID:    uint(item.ItemID),
		URL:         item.ItemURL,
		Text:        item.ItemText,
		Title:       item.ItemTitle,
		By:          item.ItemBy,
		ItemType:    item.ItemType,
		DescenDants: int(item.DescenDants),
		Kids:        item.Kids,
		Category:    common.ITEM_CATEGORY_DEFAULT,
		Score:       int(item.ItemScore),
		CreatedTime: int64(item.ItemTime),
		UpdatedAt:   currentTime.Unix(),
		CreatedAt:   currentTime.Unix(),
		ItemStatus:  common.ITEM_STATUS_NEW,
	}
	return itemEntity
}

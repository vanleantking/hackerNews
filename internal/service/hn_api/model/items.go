package model

import (
	"hackerNewsApi/internal/entity"
	"time"
)

type Item struct {
	ItemBy    string `json:"by"`
	ItemID    int    `json:"id"`
	ItemScore int    `json:"score"`
	ItemText  string `json:"text"`
	ItemTime  uint16 `json:time"`
	ItemTitle string `json:"title"`
	ItemType  string `json:"type"`
	ItemURL   string `json:"url"`
}

type HNItems struct {
	Items []int64
}

func MapperItemsCreateEntity(items []int64) *[]entity.Item {
	currentTime := time.Now()
	var tmp = make([]entity.Item, 0)
	for _, item := range items {
		itemEntity := entity.Item{
			HNItemID:  int(item),
			UpdatedAt: currentTime.Unix(),
			CreatedAt: currentTime,
		}

		tmp = append(tmp, itemEntity)
	}
	if len(tmp) > 0 {
		result := make([]entity.Item, len(tmp))
		copy(result, tmp)
		return &result
	}
	return nil
}

func MapperItemsUpsertEntity(items []Item) *[]entity.Item {
	currentTime := time.Now()
	var tmp = make([]entity.Item, 0)
	for _, item := range items {
		itemEntity := entity.Item{}
		itemEntity.HNItemID = item.ItemID
		itemEntity.URL = item.ItemURL
		itemEntity.Text = item.ItemText
		itemEntity.Title = item.ItemTitle
		itemEntity.By = item.ItemBy
		itemEntity.ItemType = item.ItemType
		itemEntity.Category = ""
		itemEntity.Score = int8(item.ItemScore)
		itemEntity.UpdatedAt = currentTime.Unix()
		itemEntity.CreatedAt = currentTime
		tmp = append(tmp, itemEntity)
	}
	if len(tmp) > 0 {
		result := make([]entity.Item, len(tmp))
		copy(result, tmp)
		return &result
	}
	return nil
}

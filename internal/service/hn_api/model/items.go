package model

import (
	"fmt"
	"hackerNewsApi/internal/entity"
	"time"
)

type Item struct {
	ItemBy    string `json:"by"`
	ItemID    uint   `json:"id"`
	ItemScore int    `json:"score"`
	ItemText  string `json:"text"`
	ItemTime  uint   `json:time"`
	ItemTitle string `json:"title"`
	ItemType  string `json:"type"`
	ItemURL   string `json:"url"`
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
			UpdatedAt: currentTime.Unix(),
			CreatedAt: currentTime.Unix(),
		}

		fmt.Println("MapperItemsCreateEntity, ", item, itemEntity.HNItemID, itemEntity.ID)

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
		itemEntity.Category = 0
		itemEntity.Score = int8(item.ItemScore)
		itemEntity.CreatedTime = int64(item.ItemTime)
		itemEntity.UpdatedAt = currentTime.Unix()
		itemEntity.CreatedAt = currentTime.Unix()
		itemEntity.ItemStatus = 0
		tmp = append(tmp, itemEntity)
	}
	if len(tmp) > 0 {
		result := make([]entity.Item, len(tmp))
		copy(result, tmp)
		return &result
	}
	return nil
}

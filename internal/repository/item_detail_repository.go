package repository

import (
	"hackerNewsApi/internal/components/logger"
	"hackerNewsApi/internal/entity"
	"hackerNewsApi/internal/service/hn_api/common"

	"gorm.io/gorm"
)

type itemDetailRepository struct {
	Repository[entity.Item]
	Log *logger.Logger
}

type ItemDetailRepository interface {
	UpdateItem(items entity.Item) error
}

func NewItemDetailRepository(log *logger.Logger, db *gorm.DB) ItemDetailRepository {
	return &itemDetailRepository{
		Repository: *NewRepository[entity.Item](db),
	}
}

func (listItems *itemDetailRepository) UpdateItem(item entity.Item) error {
	tx := listItems.DB.
		Where("hn_item_id=?", item.HNItemID).
		Assign(entity.Item{
			Score:       item.Score,
			By:          item.By,
			Title:       item.Title,
			URL:         item.URL,
			DescenDants: item.DescenDants,
			Kids:        item.Kids,
			ItemType:    item.ItemType,
			CreatedTime: item.CreatedTime,
			ItemDeleted: item.ItemDeleted,
			ItemStatus:  common.ITEM_STATUS_PROCESS_TITLE,
			UpdatedAt:   item.UpdatedAt}).
		FirstOrCreate(&item)
	return tx.Error
}

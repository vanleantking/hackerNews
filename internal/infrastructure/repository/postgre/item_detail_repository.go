package postgre

import (
	"hackerNewsApi/internal/domains/entity"
	"hackerNewsApi/internal/domains/usecases"
	"hackerNewsApi/internal/infrastructure/repository"
	"hackerNewsApi/internal/infrastructure/service/hn_api/common"
	"hackerNewsApi/pkg/logger"

	"gorm.io/gorm"
)

type itemDetailRepository struct {
	repository.Repository[entity.Item]
	Log *logger.Logger
}

func NewItemDetailRepository(log *logger.Logger, db *gorm.DB) usecases.ItemDetailRepository {
	return &itemDetailRepository{
		Repository: *repository.NewRepository[entity.Item](db),
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

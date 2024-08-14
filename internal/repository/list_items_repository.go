package repository

import (
	"hackerNewsApi/internal/components/logger"
	"hackerNewsApi/internal/entity"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type listItemRepository struct {
	Repository[entity.Item]
	Log *logger.Logger
}

type ListItemRepository interface {
	UpsertListItems(items []entity.Item) []error
	UpsertBulkItems(items []entity.Item) error
}

func NewListItemRepository(log *logger.Logger, db *gorm.DB) ListItemRepository {
	return &listItemRepository{
		Repository: *NewRepository[entity.Item](db),
	}
}

func (listItems *listItemRepository) UpsertListItems(items []entity.Item) []error {
	var errs = make([]error, 0)
	for _, item := range items {
		tx := listItems.DB.
			Where("hn_item_id=?", item.HNItemID).
			Assign(entity.Item{Score: item.Score, UpdatedAt: item.UpdatedAt}).
			FirstOrCreate(&item)
		if tx.Error != nil {
			errs = append(errs, tx.Error)
			continue
		}
		time.Sleep(100 * time.Millisecond)
	}
	return errs
}

func (listItems *listItemRepository) UpsertBulkItems(items []entity.Item) error {
	tx := listItems.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "hn_item_id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"item_score",
			"updated_at"}),
	}).Create(&items)
	return tx.Error
}

package postgre

import (
	"hackerNewsApi/pkg/logger"
	"hackerNewsApi/internal/domains/entity"
	"hackerNewsApi/internal/domains/usecases"
	"hackerNewsApi/internal/infrastructure/repository"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type listItemRepository struct {
	repository.Repository[entity.Item]
	Log *logger.Logger
}

func NewListItemRepository(log *logger.Logger, db *gorm.DB) usecases.ListItemRepository {
	return &listItemRepository{
		Repository: *repository.NewRepository[entity.Item](db),
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

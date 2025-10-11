package restaurantstorage

import (
	"context"
	"crud-go/common"
	"crud-go/module/restaurant/model"

	"gorm.io/gorm"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *restaurantmodule.RestaurantUpdate) error {
	if err := s.db.Table(restaurantmodule.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(restaurantmodule.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("like_count",
			gorm.Expr("like_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(restaurantmodule.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("like_count",
			gorm.Expr("like_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

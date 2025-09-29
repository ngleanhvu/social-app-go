package restaurantstorage

import (
	"context"
	"crud-go/common"
	"crud-go/module/restaurant/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *restaurantmodule.RestaurantUpdate) error {
	if err := s.db.Table(restaurantmodule.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

package restaurantstorage

import (
	"context"
	"crud-go/common"
	"crud-go/module/restaurant/model"
)

func (s *sqlStore) Create(context context.Context, data *restaurantmodule.RestaurantCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	data.GenUID(common.DbTypeRestaurant)
	return nil
}

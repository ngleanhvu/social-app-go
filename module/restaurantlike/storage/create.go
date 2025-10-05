package restaurantlikestorage

import (
	"context"
	"crud-go/common"
	restaurantlikemodel "crud-go/module/restaurantlike/model"
)

func (s *sqlStore) Create(ctx context.Context,
	data *restaurantlikemodel.RestaurantLikeCreate) error {
	if err := s.db.Table(restaurantlikemodel.RestaurantLike{}.
		TableName()).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

package restaurantlikestorage

import (
	"context"
	"crud-go/common"
	restaurantlikemodel "crud-go/module/restaurantlike/model"
)

func (s *sqlStore) Delete(ctx context.Context,
	data *restaurantlikemodel.RestaurantLikeUpdate) error {
	if err := s.db.Table(restaurantlikemodel.RestaurantLikeUpdate{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", data.UserId, data.RestaurantId).
		Delete(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

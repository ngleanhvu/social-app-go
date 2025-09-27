package restaurantstorage

import (
	"context"
	"crud-go/module/restaurant/model"
)

func (s *sqlStore) FindDataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodule.Restaurant, error) {
	var data restaurantmodule.Restaurant

	if err := s.db.WithContext(ctx).Where(condition).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

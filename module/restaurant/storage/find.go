package restaurantstorage

import (
	"context"
	"crud-go/common"
	"crud-go/module/restaurant/model"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodule.Restaurant, error) {
	var data restaurantmodule.Restaurant

	if err := s.db.WithContext(ctx).Where(condition).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}

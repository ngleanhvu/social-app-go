package restaurantstorage

import (
	"context"
	"crud-go/common"
	restaurantmodule "crud-go/module/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	filter *restaurantmodule.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodule.Restaurant, error) {
	var result []restaurantmodule.Restaurant

	db := s.db.Table(restaurantmodule.Restaurant{}.TableName()).Where("status = 1")

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?", f.OwnerId)
		}

		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
	}

	var total int64

	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	paging.Total = int(total)

	offset := (paging.Page - 1) * paging.Limit

	if err := db.
		Limit(paging.Limit).
		Offset(offset).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

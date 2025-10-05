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
		return nil, common.ErrDB(err)
	}

	paging.Total = int(total)

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset).Limit(paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 && len(result) < paging.Total {
		last := result[len(result)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}

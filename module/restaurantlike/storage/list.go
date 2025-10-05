package restaurantlikestorage

import (
	"context"
	"crud-go/common"
	restaurantlikemodel "crud-go/module/restaurantlike/model"
	"fmt"
	"log"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

const timeLayout = "2006-01-02T15:04:05.9999999"

func (s *sqlStore) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	if len(ids) == 0 {
		return map[int]int{}, nil
	}

	type result struct {
		RestaurantId int
		Count        int
	}

	var results []result

	if err := s.db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).
		Where("restaurant_id in ?", ids).
		Select("restaurant_id, COUNT(*) as count").
		Group("restaurant_id").
		Scan(&results).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	likeMap := make(map[int]int)

	for _, v := range results {
		likeMap[v.RestaurantId] = v.Count
	}

	return likeMap, nil
}

func (s *sqlStore) GetUsersLikeRestaurant(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]common.SimpleUser, error) {
	var results []restaurantlikemodel.RestaurantLike

	db := s.db

	db = db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}

	var total int64

	if err := db.Count(&total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	paging.Total = int(total)

	db = db.Preload("User")

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&results).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, v := range results {
		log.Println(v.RestaurantId, v.UserId, v.User.Id)
	}

	users := make([]common.SimpleUser, len(results))

	for i, item := range results {
		users[i] = *results[i].User
		log.Println("Created at", item.CreatedAt)
		users[i].CreatedAt = item.CreatedAt
		users[i].UpdatedAt = nil
		//users[i].Mask(false)
		if i == len(results)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return users, nil
}

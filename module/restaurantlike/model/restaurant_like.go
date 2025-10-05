package restaurantlikemodel

import (
	"crud-go/common"
	"errors"
	"time"
)

const EntityName = "RestaurantLike"

type RestaurantLike struct {
	RestaurantId    int        `json:"restaurant_id" gorm:"restaurant_id"`
	UserId          int        `json:"user_id" gorm:"user_id"`
	CreatedAt       *time.Time `json:"created_at" gorm:"column:created_at"`
	common.SQLModel `json:",inline"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (RestaurantLike) TableName() string {
	return "restaurant_likes"
}

type RestaurantLikeCreate struct {
	RestaurantId int        `json:"restaurant_id" gorm:"restaurant_id"`
	UserId       int        `json:"user_id" gorm:"user_id"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at"`
}

func (RestaurantLikeCreate) TableName() string {
	return RestaurantLike{}.TableName()
}

type RestaurantLikeUpdate struct {
	RestaurantId int        `json:"restaurant_id" gorm:"restaurant_id"`
	UserId       int        `json:"user_id" gorm:"user_id"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at"`
}

func (RestaurantLikeUpdate) TableName() string {
	return RestaurantLike{}.TableName()
}

var (
	ErrCannonLikeRestaurant = common.NewCustomErrorResponse(errors.New("Can not like restaurant"),
		"Can not like restaurant",
		"Can not like restaurant",
	)

	ErrCannonDislikeRestaurant = common.NewCustomErrorResponse(errors.New("Can not dislike restaurant"),
		"Can not dislike restaurant",
		"Can not dislike restaurant",
	)
)

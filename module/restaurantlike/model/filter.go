package restaurantlikemodel

type Filter struct {
	RestaurantId int `json:"restaurant_id" gorm:"column:restaurant_id"`
}

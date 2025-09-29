package restaurantmodule

import (
	"crud-go/common"
	"errors"
	"strings"
)

type RestaurantType string

const TypeNormal RestaurantType = "normal"
const TypePremium RestaurantType = "premium"
const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name"`
	Addr            string         `json:"addr" gorm:"column:addr"`
	Type            RestaurantType `json:"type" gorm:"column:type"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
	Addr            string `json:"addr" gorm:"column:addr"`
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:" column:addr"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (data *RestaurantUpdate) Validate() error {
	*data.Name = strings.TrimSpace(*data.Name)

	if *data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

var (
	ErrNameIsEmpty = errors.New(`Name cannot be empty`)
)

package usermodel

import (
	"crud-go/common"
	"errors"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email"`
	Password        string `json:"-" gorm:"column:password"`
	FirstName       string `json:"first_name" gorm:"column:first_name"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	Salt            string `json:"-" gorm:"column:salt"`
	Role            string `json:"role" gorm:"column:role"`
	Phone           string `json:"phone" gorm:"column:phone"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email"`
	Password        string `json:"password" gorm:"column:password"`
	FirstName       string `json:"first_name" gorm:"column:first_name"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	Salt            string `json:"-" gorm:"column:salt"`
	Role            string `json:"-" gorm:"column:role"`
	Phone           string `json:"phone" gorm:"column:phone"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email"`
	Password string `json:"password" form:"password" gorm:"column:password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

var (
	ErrEmailExisted = common.NewCustomErrorResponse(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
	ErrInvalidEmailOrPassword = common.NewCustomErrorResponse(
		errors.New("email or password wrong"),
		"email or password wrong",
		"email or password wrong",
	)
)

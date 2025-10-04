package userbiz

import (
	"context"
	"crud-go/common"
	usermodel "crud-go/module/user/model"
)

type RegisterStore interface {
	Find(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	Create(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}
type registerBiz struct {
	store RegisterStore
	hash  Hasher
}

func NewRegisterBiz(store RegisterStore, hash Hasher) *registerBiz {
	return &registerBiz{store: store, hash: hash}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := biz.store.Find(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)
	data.Password = biz.hash.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}

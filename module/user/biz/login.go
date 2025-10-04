package userbiz

import (
	"context"
	"crud-go/common"
	"crud-go/component/tokenprovider"
	usermodel "crud-go/module/user/model"
)

type LoginStorage interface {
	Find(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
}

type loginBiz struct {
	store         LoginStorage
	tokenprovider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(store LoginStorage,
	tokenprovider tokenprovider.Provider,
	hasher Hasher, expiry int) *loginBiz {
	return &loginBiz{store: store,
		tokenprovider: tokenprovider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.store.Find(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrInvalidEmailOrPassword
	}

	passwordHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passwordHashed {
		return nil, usermodel.ErrInvalidEmailOrPassword
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenprovider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil

}

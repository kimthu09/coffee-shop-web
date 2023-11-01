package userbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

type RegisterStorage interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	store  RegisterStorage
	hasher Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		store:  registerStorage,
		hasher: hasher,
	}
}

func (biz *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt

	id, err := common.GenerateId()
	if err != nil {
		return common.ErrInternal(err)
	}

	data.Id = id

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(common.TableUser, err)
	}

	return nil
}

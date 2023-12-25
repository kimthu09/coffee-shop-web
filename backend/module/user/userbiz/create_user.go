package userbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/component/hasher"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

type CreateUserRepo interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type createUserBiz struct {
	gen       generator.IdGenerator
	repo      CreateUserRepo
	hasher    hasher.Hasher
	requester middleware.Requester
}

func NewCreateUserBiz(
	gen generator.IdGenerator,
	repo CreateUserRepo,
	hasher hasher.Hasher,
	requester middleware.Requester) *createUserBiz {
	return &createUserBiz{
		gen:       gen,
		repo:      repo,
		hasher:    hasher,
		requester: requester,
	}
}

func (biz *createUserBiz) CreateUser(
	ctx context.Context,
	data *usermodel.UserCreate) error {
	if biz.requester.GetRoleId() != common.RoleAdminId {
		return usermodel.ErrUserCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	salt := common.GenSalt(50)
	data.Password = biz.hasher.Hash(common.DefaultPass + salt)
	data.Salt = salt

	if err := handleUserId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}

func handleUserId(gen generator.IdGenerator, data *usermodel.UserCreate) error {
	id, err := gen.GenerateId()
	if err != nil {
		return err
	}

	data.Id = id
	return nil
}

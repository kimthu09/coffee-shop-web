package userbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/component/hasher"
	"coffee_shop_management_backend/component/token_provider"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

type LoginRepo interface {
	FindUserByEmail(
		ctx context.Context,
		email string,
	) (*usermodel.User, error)
}

type loginBiz struct {
	appCtx        appctx.AppContext
	repo          LoginRepo
	expiry        int
	tokenProvider token_provider.Provider
	hasher        hasher.Hasher
}

func NewLoginBiz(
	appCtx appctx.AppContext,
	repo LoginRepo,
	expiry int,
	tokenProvider token_provider.Provider,
	hasher hasher.Hasher) *loginBiz {
	return &loginBiz{
		appCtx:        appCtx,
		repo:          repo,
		expiry:        expiry,
		tokenProvider: tokenProvider,
		hasher:        hasher,
	}
}

// 1. Find user, email
// 2. Hash pass from input & compare with pass in db
// 3. Provider: issue JWT token for Client
// 3.1 Access token & Refresh token
// 4. Return token(s)

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := biz.repo.FindUserByEmail(ctx, data.Email)

	if err != nil {
		return nil, usermodel.ErrUserEmailOrPasswordInvalid
	}

	passwordHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passwordHashed {
		return nil, usermodel.ErrUserEmailOrPasswordInvalid
	}

	payload := token_provider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role.Id,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}

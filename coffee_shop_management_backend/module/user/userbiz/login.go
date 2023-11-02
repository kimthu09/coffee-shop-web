package userbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/component/token_provider"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

type LoginStorage interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
}

type loginBusiness struct {
	appCtx        appctx.AppContext
	userStore     LoginStorage
	expiry        int
	tokenProvider token_provider.Provider
	hasher        Hasher
}

func NewLoginBusiness(appCtx appctx.AppContext,
	userStore LoginStorage,
	expiry int,
	tokenProvider token_provider.Provider,
	hasher Hasher) *loginBusiness {
	return &loginBusiness{
		appCtx:        appCtx,
		userStore:     userStore,
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

func (biz *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := biz.userStore.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	passwordHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passwordHashed {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	payload := token_provider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
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

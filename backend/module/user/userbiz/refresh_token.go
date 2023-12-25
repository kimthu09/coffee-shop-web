package userbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/tokenprovider"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

type refreshTokenBiz struct {
	expiry        int
	tokenProvider tokenprovider.Provider
}

func NewRefreshTokenBiz(
	expiry int,
	tokenProvider tokenprovider.Provider) *refreshTokenBiz {
	return &refreshTokenBiz{
		expiry:        expiry,
		tokenProvider: tokenProvider,
	}
}

func (biz *refreshTokenBiz) RefreshToken(
	ctx context.Context,
	refreshToken *usermodel.UserRefreshToken) (*usermodel.Account, error) {
	payload, err := biz.tokenProvider.Validate(refreshToken.RefreshToken)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	accessToken, err := biz.tokenProvider.Generate(*payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken, nil)

	return account, nil
}

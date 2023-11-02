package usermodel

import (
	"coffee_shop_management_backend/component/token_provider"
)

type Account struct {
	AccessToken  *token_provider.Token `json:"access_token"`
	RefreshToken *token_provider.Token `json:"refresh_token"`
}

func NewAccount(at, rt *token_provider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

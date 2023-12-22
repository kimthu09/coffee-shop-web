package tokenprovider

import (
	"coffee_shop_management_backend/common"
	"errors"
	"time"
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
}

var (
	ErrInvalidToken = common.NewCustomError(errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)
)

package userbiz

import (
	"coffee_shop_management_backend/component/tokenprovider"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestNewRefreshTokenBiz(t *testing.T) {
	type args struct {
		expiry        int
		tokenProvider tokenprovider.Provider
	}
	tests := []struct {
		name string
		args args
		want *refreshTokenBiz
	}{
		{
			name: "Create object has type refreshTokenBiz",
			args: args{
				expiry:        0,
				tokenProvider: nil,
			},
			want: &refreshTokenBiz{
				expiry:        0,
				tokenProvider: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRefreshTokenBiz(
				tt.args.expiry,
				tt.args.tokenProvider,
			)

			assert.Equal(t, tt.want, got, "NewLoginBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_refreshTokenBiz_RefreshToken(t *testing.T) {
	type fields struct {
		expiry        int
		tokenProvider tokenprovider.Provider
	}
	type args struct {
		ctx          context.Context
		refreshToken *usermodel.UserRefreshToken
	}

	mockTokenPro := new(mockTokenProvider)

	userId := mock.Anything
	roleId := mock.Anything
	expiry := 1000
	refreshToken := mock.Anything
	payload := tokenprovider.TokenPayload{
		UserId: userId,
		Role:   roleId,
	}
	token := tokenprovider.Token{
		Token:   mock.Anything,
		Created: time.Time{},
		Expiry:  expiry,
	}

	account := usermodel.Account{
		AccessToken:  &token,
		RefreshToken: nil,
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *usermodel.Account
		wantErr bool
	}{
		{
			name: "Refresh token failed because refreshToken is invalid",
			fields: fields{
				expiry:        expiry,
				tokenProvider: mockTokenPro,
			},
			args: args{
				ctx:          context.Background(),
				refreshToken: &usermodel.UserRefreshToken{RefreshToken: refreshToken},
			},
			mock: func() {
				mockTokenPro.
					On(
						"Validate",
						refreshToken).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Refresh token failed because can not generate token",
			fields: fields{
				expiry:        expiry,
				tokenProvider: mockTokenPro,
			},
			args: args{
				ctx:          context.Background(),
				refreshToken: &usermodel.UserRefreshToken{RefreshToken: refreshToken},
			},
			mock: func() {
				mockTokenPro.
					On(
						"Validate",
						refreshToken).
					Return(&payload, nil).
					Once()

				mockTokenPro.
					On(
						"Generate",
						payload,
						expiry).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Refresh token successfully",
			fields: fields{
				expiry:        expiry,
				tokenProvider: mockTokenPro,
			},
			args: args{
				ctx:          context.Background(),
				refreshToken: &usermodel.UserRefreshToken{RefreshToken: refreshToken},
			},
			mock: func() {
				mockTokenPro.
					On(
						"Validate",
						refreshToken).
					Return(&payload, nil).
					Once()

				mockTokenPro.
					On(
						"Generate",
						payload,
						expiry).
					Return(&token, nil).
					Once()
			},
			want:    &account,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &refreshTokenBiz{
				expiry:        tt.fields.expiry,
				tokenProvider: tt.fields.tokenProvider,
			}

			tt.mock()

			got, err := biz.RefreshToken(tt.args.ctx, tt.args.refreshToken)

			if tt.wantErr {
				assert.NotNil(t, err, "RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "RefreshToken = %v, want = %v", got, tt.want)
			}
		})
	}
}

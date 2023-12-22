package userbiz

import (
	"coffee_shop_management_backend/component/hasher"
	"coffee_shop_management_backend/component/tokenprovider"
	"coffee_shop_management_backend/module/role/rolemodel"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type mockLoginRepo struct {
	mock.Mock
}

func (m *mockLoginRepo) FindUserByEmail(
	ctx context.Context,
	email string) (*usermodel.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usermodel.User), args.Error(1)
}

type mockTokenProvider struct {
	mock.Mock
}

func (m *mockTokenProvider) Generate(
	data tokenprovider.TokenPayload,
	expiry int) (*tokenprovider.Token, error) {
	args := m.Called(data, expiry)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*tokenprovider.Token), args.Error(1)
}

func (m *mockTokenProvider) Validate(
	token string) (*tokenprovider.TokenPayload, error) {
	args := m.Called(token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*tokenprovider.TokenPayload), args.Error(1)
}

func TestNewLoginBiz(t *testing.T) {
	type args struct {
		repo          LoginRepo
		tokenProvider tokenprovider.Provider
		hasher        hasher.Hasher
		expiry        int
		expiryRefresh int
	}

	mockRepo := new(mockLoginRepo)
	mockHash := new(mockHasher)
	mockTokenPro := new(mockTokenProvider)

	tests := []struct {
		name string
		args args
		want *loginBiz
	}{
		{
			name: "Create object has type LoginBiz",
			args: args{
				repo:          mockRepo,
				tokenProvider: mockTokenPro,
				hasher:        mockHash,
				expiry:        1000,
				expiryRefresh: 1000,
			},
			want: &loginBiz{
				repo:          mockRepo,
				expiry:        1000,
				expiryRefresh: 1000,
				tokenProvider: mockTokenPro,
				hasher:        mockHash,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLoginBiz(
				tt.args.repo,
				tt.args.expiry,
				tt.args.expiryRefresh,
				tt.args.tokenProvider,
				tt.args.hasher,
			)

			assert.Equal(t, tt.want, got, "NewLoginBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_loginBiz_Login(t *testing.T) {
	type fields struct {
		repo          LoginRepo
		expiry        int
		expiryRefresh int
		tokenProvider tokenprovider.Provider
		hasher        hasher.Hasher
	}
	type args struct {
		ctx  context.Context
		data *usermodel.UserLogin
	}

	mockRepo := new(mockLoginRepo)
	mockHash := new(mockHasher)
	mockTokenPro := new(mockTokenProvider)
	expiry := 1000
	expiryRefresh := 2000
	requestEmail := "a@gmail.com"
	requestPassword := "123456"
	salt := mock.Anything
	hashedPass := mock.Anything
	userId := mock.Anything
	roleId := mock.Anything
	role := rolemodel.Role{Id: roleId}
	userLogin := usermodel.UserLogin{
		Email:    requestEmail,
		Password: requestPassword,
	}
	user := usermodel.User{
		Id:       userId,
		Email:    requestEmail,
		Password: hashedPass,
		Salt:     salt,
		RoleId:   roleId,
		Role:     role,
	}
	payload := tokenprovider.TokenPayload{
		UserId: userId,
		Role:   roleId,
	}
	token := tokenprovider.Token{
		Token:   mock.Anything,
		Created: time.Time{},
		Expiry:  expiry,
	}
	refreshToken := tokenprovider.Token{
		Token:   mock.Anything,
		Created: time.Time{},
		Expiry:  expiryRefresh,
	}
	account := usermodel.Account{
		AccessToken:  &token,
		RefreshToken: &refreshToken,
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
			name: "Login failed because can not found user has same email",
			fields: fields{
				repo:          mockRepo,
				expiry:        expiry,
				expiryRefresh: expiryRefresh,
				tokenProvider: mockTokenPro,
				hasher:        mockHash,
			},
			args: args{
				ctx:  context.Background(),
				data: &userLogin,
			},
			mock: func() {
				mockRepo.
					On(
						"FindUserByEmail",
						context.Background(),
						requestEmail).
					Return(nil, mockErr).
					Once()
			},
			want:    &account,
			wantErr: true,
		},
		{
			name: "Login failed because wrong password",
			fields: fields{
				repo:          mockRepo,
				expiry:        expiry,
				expiryRefresh: expiryRefresh,
				tokenProvider: mockTokenPro,
				hasher:        mockHash,
			},
			args: args{
				ctx:  context.Background(),
				data: &userLogin,
			},
			mock: func() {
				mockRepo.
					On(
						"FindUserByEmail",
						context.Background(),
						requestEmail).
					Return(&user, nil).
					Once()

				mockHash.
					On(
						"Hash",
						requestPassword+salt).
					Return("This is wrong pass").
					Once()
			},
			want:    &account,
			wantErr: true,
		},
		{
			name: "Login failed because can not generate accessToken",
			fields: fields{
				repo:          mockRepo,
				expiry:        expiry,
				expiryRefresh: expiryRefresh,
				tokenProvider: mockTokenPro,
				hasher:        mockHash,
			},
			args: args{
				ctx:  context.Background(),
				data: &userLogin,
			},
			mock: func() {
				mockRepo.
					On(
						"FindUserByEmail",
						context.Background(),
						requestEmail).
					Return(&user, nil).
					Once()

				mockHash.
					On(
						"Hash",
						requestPassword+salt).
					Return(hashedPass).
					Once()

				mockTokenPro.
					On(
						"Generate",
						payload,
						expiry,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    &account,
			wantErr: true,
		},
		{
			name: "Login failed because can not generate refreshToken",
			fields: fields{
				repo:          mockRepo,
				expiry:        expiry,
				expiryRefresh: expiryRefresh,
				tokenProvider: mockTokenPro,
				hasher:        mockHash,
			},
			args: args{
				ctx:  context.Background(),
				data: &userLogin,
			},
			mock: func() {
				mockRepo.
					On(
						"FindUserByEmail",
						context.Background(),
						requestEmail).
					Return(&user, nil).
					Once()

				mockHash.
					On(
						"Hash",
						requestPassword+salt).
					Return(hashedPass).
					Once()

				mockTokenPro.
					On(
						"Generate",
						payload,
						expiry,
					).
					Return(&token, nil).
					Once()

				mockTokenPro.
					On(
						"Generate",
						payload,
						expiryRefresh,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    &account,
			wantErr: true,
		},
		{
			name: "Login successfully",
			fields: fields{
				repo:          mockRepo,
				expiry:        expiry,
				expiryRefresh: expiryRefresh,
				tokenProvider: mockTokenPro,
				hasher:        mockHash,
			},
			args: args{
				ctx:  context.Background(),
				data: &userLogin,
			},
			mock: func() {
				mockRepo.
					On(
						"FindUserByEmail",
						context.Background(),
						requestEmail).
					Return(&user, nil).
					Once()

				mockHash.
					On(
						"Hash",
						requestPassword+salt).
					Return(hashedPass).
					Once()

				mockTokenPro.
					On(
						"Generate",
						payload,
						expiry,
					).
					Return(&token, nil).
					Once()

				mockTokenPro.
					On(
						"Generate",
						payload,
						expiryRefresh,
					).
					Return(&refreshToken, nil).
					Once()
			},
			want:    &account,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &loginBiz{
				repo:          tt.fields.repo,
				expiry:        tt.fields.expiry,
				expiryRefresh: tt.fields.expiryRefresh,
				tokenProvider: tt.fields.tokenProvider,
				hasher:        tt.fields.hasher,
			}

			tt.mock()

			got, err := biz.Login(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "Login() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "Login() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "Login = %v, want = %v", got, tt.want)
			}
		})
	}
}

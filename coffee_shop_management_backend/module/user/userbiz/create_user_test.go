package userbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/component/hasher"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/role/rolemodel"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockCreateUserRepo struct {
	mock.Mock
}

func (m *mockCreateUserRepo) CreateUser(
	ctx context.Context,
	data *usermodel.UserCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockIdGenerator struct {
	mock.Mock
}

func (m *mockIdGenerator) GenerateId() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *mockIdGenerator) IdProcess(id *string) (*string, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*string), args.Error(1)
}

type mockHasher struct {
	mock.Mock
}

func (m *mockHasher) Hash(data string) string {
	args := m.Called(data)
	return args.Get(0).(string)
}

func TestNewCreateUserBiz(t *testing.T) {
	type args struct {
		gen       generator.IdGenerator
		repo      CreateUserRepo
		hasher    hasher.Hasher
		requester middleware.Requester
	}

	mockGenerator := new(mockIdGenerator)
	mockRepo := new(mockCreateUserRepo)
	mockHash := new(mockHasher)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *createUserBiz
	}{
		{
			name: "Create object has type CreateUserBiz",
			args: args{
				gen:       mockGenerator,
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
			want: &createUserBiz{
				gen:       mockGenerator,
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateUserBiz(
				tt.args.gen,
				tt.args.repo,
				tt.args.hasher,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewCreateUserBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createUserBiz_CreateUser(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		repo      CreateUserRepo
		hasher    hasher.Hasher
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data *usermodel.UserCreate
	}

	mockGenerator := new(mockIdGenerator)
	mockRepo := new(mockCreateUserRepo)
	mockHash := new(mockHasher)
	mockRequest := new(mockRequester)

	validId := "012345678901"
	validEmail := "a@gmail.com"

	salt := ""
	for i := 0; i <= 50; i++ {
		salt = salt + "a"
	}

	password := "hashedPassword"

	data := usermodel.UserCreate{
		Name:     validId,
		Email:    validEmail,
		RoleId:   validId,
		Password: password,
		Salt:     salt,
	}
	adminRole := rolemodel.Role{Id: common.RoleAdminId}
	notAdminRole := rolemodel.Role{Id: mock.Anything}

	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create user failed because user is not allowed",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &data,
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(notAdminRole.Id).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create user failed because data is invalid",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
				data: &usermodel.UserCreate{
					Name:   "",
					Email:  "",
					RoleId: "",
				},
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(adminRole.Id).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create user failed because can not generate id",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &data,
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(adminRole.Id).
					Once()

				mockHash.
					On(
						"Hash",
						mock.Anything,
					).
					Return(password).
					Once()

				mockGenerator.
					On(
						"GenerateId",
					).
					Return(mock.Anything, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create user failed because can not store to database",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &data,
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(adminRole.Id).
					Once()

				mockHash.
					On(
						"Hash",
						mock.Anything,
					).
					Return(password).
					Once()

				mockGenerator.
					On(
						"GenerateId",
					).
					Return(validId, nil).
					Once()

				mockRepo.
					On(
						"CreateUser",
						context.Background(),
						&data,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create user successfully",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &data,
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(adminRole.Id).
					Once()

				mockHash.
					On(
						"Hash",
						mock.Anything,
					).
					Return(password).
					Once()

				mockGenerator.
					On(
						"GenerateId",
					).
					Return(validId, nil).
					Once()

				mockRepo.
					On(
						"CreateUser",
						context.Background(),
						&data,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &createUserBiz{
				gen:       tt.fields.gen,
				repo:      tt.fields.repo,
				hasher:    tt.fields.hasher,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.CreateUser(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

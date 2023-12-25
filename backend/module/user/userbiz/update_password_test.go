package userbiz

import (
	"coffee_shop_management_backend/component/hasher"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockUpdatePassUser struct {
	mock.Mock
}

func (m *mockUpdatePassUser) GetUser(
	ctx context.Context,
	userId string) (*usermodel.User, error) {
	args := m.Called(ctx, userId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usermodel.User), args.Error(1)
}
func (m *mockUpdatePassUser) UpdateUserPassword(
	ctx context.Context,
	id string,
	pass string) error {
	args := m.Called(ctx, id, pass)
	return args.Error(0)
}

func TestNewUpdatePasswordBiz(t *testing.T) {
	type args struct {
		repo   UpdatePasswordRepo
		hasher hasher.Hasher
	}

	mockRepo := new(mockUpdatePassUser)
	mockHash := new(mockHasher)

	tests := []struct {
		name string
		args args
		want *updatePasswordBiz
	}{
		{
			name: "Create object has type UpdatePassUserRepo",
			args: args{
				repo:   mockRepo,
				hasher: mockHash,
			},
			want: &updatePasswordBiz{
				repo:   mockRepo,
				hasher: mockHash,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdatePasswordBiz(
				tt.args.repo,
				tt.args.hasher,
			)

			assert.Equal(t, tt.want, got, "NewUpdatePasswordBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_updatePasswordBiz_UpdatePassword(t *testing.T) {
	type fields struct {
		repo      UpdatePasswordRepo
		hasher    hasher.Hasher
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		id   string
		data *usermodel.UserUpdatePassword
	}

	mockRepo := new(mockUpdatePassUser)
	mockHash := new(mockHasher)

	userId := mock.Anything
	oldPassword := mock.Anything
	newPassword := mock.Anything
	hashedPassword := mock.Anything
	hashedNewPassword := mock.Anything

	userUpdatePass := usermodel.UserUpdatePassword{
		OldPassword: oldPassword,
		NewPassword: newPassword,
	}
	inActiveUser := usermodel.User{
		Id:       userId,
		Password: hashedPassword,
		Salt:     mock.Anything,
		IsActive: false,
	}
	user := usermodel.User{
		Id:       userId,
		Password: hashedPassword,
		Salt:     mock.Anything,
		IsActive: true,
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update password user failed because data is invalid",
			fields: fields{
				repo:   mockRepo,
				hasher: mockHash,
			},
			args: args{
				ctx: context.Background(),
				id:  userId,
				data: &usermodel.UserUpdatePassword{
					OldPassword: "",
				},
			},
			mock: func() {

			},
			wantErr: true,
		},
		{
			name: "Update password user failed because can not get user",
			fields: fields{
				repo:   mockRepo,
				hasher: mockHash,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &userUpdatePass,
			},
			mock: func() {
				mockRepo.
					On(
						"GetUser",
						context.Background(),
						userId).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update password user failed because user is inactive",
			fields: fields{
				repo:   mockRepo,
				hasher: mockHash,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &userUpdatePass,
			},
			mock: func() {
				mockRepo.
					On(
						"GetUser",
						context.Background(),
						userId).
					Return(&inActiveUser, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update password user failed because password is wrong",
			fields: fields{
				repo:   mockRepo,
				hasher: mockHash,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &userUpdatePass,
			},
			mock: func() {
				mockRepo.
					On(
						"GetUser",
						context.Background(),
						userId).
					Return(&user, nil).
					Once()

				mockHash.
					On(
						"Hash",
						oldPassword+user.Salt).
					Return("").
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update password user failed because can not save to database",
			fields: fields{
				repo:   mockRepo,
				hasher: mockHash,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &userUpdatePass,
			},
			mock: func() {
				mockRepo.
					On(
						"GetUser",
						context.Background(),
						userId).
					Return(&user, nil).
					Once()

				mockHash.
					On(
						"Hash",
						oldPassword+user.Salt).
					Return(hashedPassword).
					Once()

				mockHash.
					On(
						"Hash",
						newPassword+user.Salt).
					Return(hashedNewPassword).
					Once()

				mockRepo.
					On(
						"UpdateUserPassword",
						context.Background(),
						userId,
						hashedNewPassword).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update password user successfully",
			fields: fields{
				repo:   mockRepo,
				hasher: mockHash,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &userUpdatePass,
			},
			mock: func() {
				mockRepo.
					On(
						"GetUser",
						context.Background(),
						userId).
					Return(&user, nil).
					Once()

				mockHash.
					On(
						"Hash",
						oldPassword+user.Salt).
					Return(hashedPassword).
					Once()

				mockHash.
					On(
						"Hash",
						newPassword+user.Salt).
					Return(hashedNewPassword).
					Once()

				mockRepo.
					On(
						"UpdateUserPassword",
						context.Background(),
						userId,
						hashedNewPassword).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &updatePasswordBiz{
				repo:      tt.fields.repo,
				hasher:    tt.fields.hasher,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.UpdatePassword(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdatePassword() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

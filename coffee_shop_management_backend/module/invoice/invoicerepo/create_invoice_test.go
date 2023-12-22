package invoicerepo

import (
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailmodel"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/shopgeneral/shopgeneralmodel"
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockInvoiceStore struct {
	mock.Mock
}

func (m *mockInvoiceStore) CreateInvoice(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockInvoiceDetailStore struct {
	mock.Mock
}

func (m *mockInvoiceDetailStore) CreateListImportNoteDetail(ctx context.Context, data []invoicedetailmodel.InvoiceDetailCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockCustomerStore struct {
	mock.Mock
}

func (m *mockCustomerStore) FindCustomer(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*customermodel.Customer, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*customermodel.Customer), args.Error(1)
}
func (m *mockCustomerStore) UpdateCustomerPoint(ctx context.Context, id string, data *customermodel.CustomerUpdatePoint) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

type mockSizeFoodStore struct {
	mock.Mock
}

func (m *mockSizeFoodStore) FindSizeFood(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*sizefoodmodel.SizeFood, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sizefoodmodel.SizeFood), args.Error(1)
}

type mockFoodStore struct {
	mock.Mock
}

func (m *mockFoodStore) FindFood(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*productmodel.Food, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*productmodel.Food), args.Error(1)
}

type mockToppingStore struct {
	mock.Mock
}

func (m *mockToppingStore) FindTopping(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*productmodel.Topping, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*productmodel.Topping), args.Error(1)
}

type mockIngredientStore struct {
	mock.Mock
}

func (m *mockIngredientStore) FindIngredient(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*ingredientmodel.Ingredient, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*ingredientmodel.Ingredient), args.Error(1)
}
func (m *mockIngredientStore) UpdateAmountIngredient(ctx context.Context, id string, data *ingredientmodel.IngredientUpdateAmount) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

type mockShopGeneralStore struct {
	mock.Mock
}

func (m *mockShopGeneralStore) FindShopGeneral(ctx context.Context) (*shopgeneralmodel.ShopGeneral, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*shopgeneralmodel.ShopGeneral), args.Error(1)
}

func TestNewCreateInvoiceRepo(t *testing.T) {
	type args struct {
		invoiceStore       InvoiceStore
		invoiceDetailStore InvoiceDetailStore
		customerStore      CustomerStore
		sizeFoodStore      SizeFoodStore
		foodStore          FoodStore
		toppingStore       ToppingStore
		ingredientStore    IngredientStore
		shopGeneralStore   ShopGeneralStore
	}

	invoiceStore := new(mockInvoiceStore)
	invoiceDetailStore := new(mockInvoiceDetailStore)
	customerStore := new(mockCustomerStore)
	sizeFoodStore := new(mockSizeFoodStore)
	foodStore := new(mockFoodStore)
	toppingStore := new(mockToppingStore)
	ingredientStore := new(mockIngredientStore)
	shopGeneralStore := new(mockShopGeneralStore)

	tests := []struct {
		name string
		args args
		want *createInvoiceRepo
	}{
		{
			name: "Create object has type createInvoiceRepo",
			args: args{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			want: &createInvoiceRepo{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateInvoiceRepo(
				tt.args.invoiceStore,
				tt.args.invoiceDetailStore,
				tt.args.customerStore,
				tt.args.sizeFoodStore,
				tt.args.foodStore,
				tt.args.toppingStore,
				tt.args.ingredientStore,
				tt.args.shopGeneralStore)
			assert.Equal(
				t, tt.want, got,
				"NewListImportNoteRepo() = %v, want %v",
				got, tt.want)
		})
	}
}

func Test_createInvoiceRepo_FindCustomer(t *testing.T) {
	type fields struct {
		invoiceStore       InvoiceStore
		invoiceDetailStore InvoiceDetailStore
		customerStore      CustomerStore
		sizeFoodStore      SizeFoodStore
		foodStore          FoodStore
		toppingStore       ToppingStore
		ingredientStore    IngredientStore
		shopGeneralStore   ShopGeneralStore
	}
	type args struct {
		ctx        context.Context
		customerId string
	}

	invoiceStore := new(mockInvoiceStore)
	invoiceDetailStore := new(mockInvoiceDetailStore)
	customerStore := new(mockCustomerStore)
	sizeFoodStore := new(mockSizeFoodStore)
	foodStore := new(mockFoodStore)
	toppingStore := new(mockToppingStore)
	ingredientStore := new(mockIngredientStore)
	shopGeneralStore := new(mockShopGeneralStore)

	customerId := "Customer001"
	customer := customermodel.Customer{
		Id: "Customer001",
	}
	var moreKeys []string
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *customermodel.Customer
		wantErr bool
	}{
		{
			name: "Find customer failed because can not get data from database",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
			},
			mock: func() {
				customerStore.
					On("FindCustomer",
						context.Background(),
						map[string]interface{}{"id": customerId},
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    &customer,
			wantErr: true,
		},
		{
			name: "Find customer successfully",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
			},
			mock: func() {
				customerStore.
					On("FindCustomer",
						context.Background(),
						map[string]interface{}{"id": customerId},
						moreKeys).
					Return(&customer, nil).
					Once()
			},
			want:    &customer,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createInvoiceRepo{
				invoiceStore:       tt.fields.invoiceStore,
				invoiceDetailStore: tt.fields.invoiceDetailStore,
				customerStore:      tt.fields.customerStore,
				sizeFoodStore:      tt.fields.sizeFoodStore,
				foodStore:          tt.fields.foodStore,
				toppingStore:       tt.fields.toppingStore,
				ingredientStore:    tt.fields.ingredientStore,
				shopGeneralStore:   tt.fields.shopGeneralStore,
			}

			tt.mock()

			got, err := repo.FindCustomer(tt.args.ctx, tt.args.customerId)

			if tt.wantErr {
				assert.NotNil(t, err, "ListInvoice() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListInvoice() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListInvoice() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_createInvoiceRepo_GetShopGeneral(t *testing.T) {
	type fields struct {
		invoiceStore       InvoiceStore
		invoiceDetailStore InvoiceDetailStore
		customerStore      CustomerStore
		sizeFoodStore      SizeFoodStore
		foodStore          FoodStore
		toppingStore       ToppingStore
		ingredientStore    IngredientStore
		shopGeneralStore   ShopGeneralStore
	}

	invoiceStore := new(mockInvoiceStore)
	invoiceDetailStore := new(mockInvoiceDetailStore)
	customerStore := new(mockCustomerStore)
	sizeFoodStore := new(mockSizeFoodStore)
	foodStore := new(mockFoodStore)
	toppingStore := new(mockToppingStore)
	ingredientStore := new(mockIngredientStore)
	shopGeneralStore := new(mockShopGeneralStore)

	mockErr := errors.New(mock.Anything)
	shopGeneral := shopgeneralmodel.ShopGeneral{
		Name:                   mock.Anything,
		Email:                  mock.Anything,
		Phone:                  mock.Anything,
		Address:                mock.Anything,
		WifiPass:               mock.Anything,
		AccumulatePointPercent: 0.001,
		UsePointPercent:        1,
	}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *shopgeneralmodel.ShopGeneral
		wantErr bool
	}{
		{
			name: "Get shop general failed because can not get data from database",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				shopGeneralStore.
					On("FindShopGeneral",
						context.Background()).
					Return(nil, mockErr).
					Once()
			},
			want:    &shopGeneral,
			wantErr: true,
		},
		{
			name: "Get shop general successfully",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				shopGeneralStore.
					On("FindShopGeneral",
						context.Background()).
					Return(&shopGeneral, nil).
					Once()
			},
			want:    &shopGeneral,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createInvoiceRepo{
				invoiceStore:       tt.fields.invoiceStore,
				invoiceDetailStore: tt.fields.invoiceDetailStore,
				customerStore:      tt.fields.customerStore,
				sizeFoodStore:      tt.fields.sizeFoodStore,
				foodStore:          tt.fields.foodStore,
				toppingStore:       tt.fields.toppingStore,
				ingredientStore:    tt.fields.ingredientStore,
				shopGeneralStore:   tt.fields.shopGeneralStore,
			}

			tt.mock()

			got, err := repo.GetShopGeneral(tt.args.ctx)

			if tt.wantErr {
				assert.NotNil(t, err, "GetShopGeneral() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "GetShopGeneral() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "GetShopGeneral() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_createInvoiceRepo_HandleCheckPermissionStatus(t *testing.T) {
	type fields struct {
		invoiceStore       InvoiceStore
		invoiceDetailStore InvoiceDetailStore
		customerStore      CustomerStore
		sizeFoodStore      SizeFoodStore
		foodStore          FoodStore
		toppingStore       ToppingStore
		ingredientStore    IngredientStore
		shopGeneralStore   ShopGeneralStore
	}
	type args struct {
		ctx  context.Context
		data *invoicemodel.InvoiceCreate
	}

	invoiceStore := new(mockInvoiceStore)
	invoiceDetailStore := new(mockInvoiceDetailStore)
	customerStore := new(mockCustomerStore)
	sizeFoodStore := new(mockSizeFoodStore)
	foodStore := new(mockFoodStore)
	toppingStore := new(mockToppingStore)
	ingredientStore := new(mockIngredientStore)
	shopGeneralStore := new(mockShopGeneralStore)

	invoiceCreate := invoicemodel.InvoiceCreate{
		InvoiceDetails: []invoicedetailmodel.InvoiceDetailCreate{
			{
				FoodId: "food001",
				Toppings: &invoicedetailmodel.InvoiceDetailToppings{
					{
						Id: "topping001",
					},
				},
			},
		},
	}
	activeFood := productmodel.Food{
		Product: &productmodel.Product{
			Id:       "food001",
			IsActive: true,
		},
	}
	inActiveFood := productmodel.Food{
		Product: &productmodel.Product{
			Id:       "food001",
			IsActive: false,
		},
	}
	activeTopping := productmodel.Topping{
		Product: &productmodel.Product{
			Id:       "topping001",
			IsActive: true,
		},
	}
	inActiveTopping := productmodel.Topping{
		Product: &productmodel.Product{
			Id:       "topping001",
			IsActive: false,
		},
	}
	var moreKeys []string
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Handle check permission status failed because can not find food",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreate,
			},
			mock: func() {
				foodStore.
					On("FindFood",
						context.Background(),
						map[string]interface{}{
							"Id": "food001",
						},
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle check permission status failed because food is inactive",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreate,
			},
			mock: func() {
				foodStore.
					On("FindFood",
						context.Background(),
						map[string]interface{}{
							"Id": "food001",
						},
						moreKeys).
					Return(&inActiveFood, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle check permission status failed because can not find topping",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreate,
			},
			mock: func() {
				foodStore.
					On("FindFood",
						context.Background(),
						map[string]interface{}{
							"Id": "food001",
						},
						moreKeys).
					Return(&activeFood, nil).
					Once()

				toppingStore.
					On("FindTopping",
						context.Background(),
						map[string]interface{}{
							"Id": "topping001",
						},
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle check permission status failed because exist inactive topping",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreate,
			},
			mock: func() {
				foodStore.
					On("FindFood",
						context.Background(),
						map[string]interface{}{
							"Id": "food001",
						},
						moreKeys).
					Return(&activeFood, nil).
					Once()

				toppingStore.
					On("FindTopping",
						context.Background(),
						map[string]interface{}{
							"Id": "topping001",
						},
						moreKeys).
					Return(&inActiveTopping, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle check permission status successfully",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreate,
			},
			mock: func() {
				foodStore.
					On("FindFood",
						context.Background(),
						map[string]interface{}{
							"Id": "food001",
						},
						moreKeys).
					Return(&activeFood, nil).
					Once()

				toppingStore.
					On("FindTopping",
						context.Background(),
						map[string]interface{}{
							"Id": "topping001",
						},
						moreKeys).
					Return(&activeTopping, nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createInvoiceRepo{
				invoiceStore:       tt.fields.invoiceStore,
				invoiceDetailStore: tt.fields.invoiceDetailStore,
				customerStore:      tt.fields.customerStore,
				sizeFoodStore:      tt.fields.sizeFoodStore,
				foodStore:          tt.fields.foodStore,
				toppingStore:       tt.fields.toppingStore,
				ingredientStore:    tt.fields.ingredientStore,
				shopGeneralStore:   tt.fields.shopGeneralStore,
			}

			tt.mock()

			err := repo.HandleCheckPermissionStatus(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "HandleCheckPermissionStatus() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleCheckPermissionStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createInvoiceRepo_HandleData(t *testing.T) {
	type fields struct {
		invoiceStore       InvoiceStore
		invoiceDetailStore InvoiceDetailStore
		customerStore      CustomerStore
		sizeFoodStore      SizeFoodStore
		foodStore          FoodStore
		toppingStore       ToppingStore
		ingredientStore    IngredientStore
		shopGeneralStore   ShopGeneralStore
	}
	type args struct {
		ctx  context.Context
		data *invoicemodel.InvoiceCreate
	}

	invoiceStore := new(mockInvoiceStore)
	invoiceDetailStore := new(mockInvoiceDetailStore)
	customerStore := new(mockCustomerStore)
	sizeFoodStore := new(mockSizeFoodStore)
	foodStore := new(mockFoodStore)
	toppingStore := new(mockToppingStore)
	ingredientStore := new(mockIngredientStore)
	shopGeneralStore := new(mockShopGeneralStore)

	measureType := enum.Weight
	sizeFood := sizefoodmodel.SizeFood{
		FoodId:   "Food001",
		SizeId:   "Size001",
		Name:     "Size Name",
		Cost:     10000,
		Price:    8000,
		RecipeId: "Recipe001",
		Recipe: recipemodel.Recipe{
			Id: "Recipe001",
			Details: recipemodel.RecipeDetails{
				{
					RecipeId:     "Recipe001",
					IngredientId: "Ingredient1",
					Ingredient: ingredientmodel.SimpleIngredient{
						Id:          "Ingredient1",
						Name:        "Ingredient Name",
						MeasureType: &measureType,
					},
					AmountNeed: 10,
				},
			},
		},
	}
	sizeFoodMoreKeys := []string{"Recipe.Details.Ingredient"}
	topping := productmodel.Topping{
		Product: &productmodel.Product{
			Id:           "Topping001",
			Name:         "Topping Name",
			Description:  mock.Anything,
			CookingGuide: mock.Anything,
			IsActive:     true,
		},
		Cost:     5000,
		Price:    4000,
		RecipeId: "Recipe002",
		Recipe: &recipemodel.Recipe{
			Id: "Recipe002",
			Details: recipemodel.RecipeDetails{
				{
					RecipeId:     "Recipe002",
					IngredientId: "Ingredient1",
					Ingredient: ingredientmodel.SimpleIngredient{
						Id:          "Ingredient1",
						Name:        "Ingredient Name",
						MeasureType: &measureType,
					},
					AmountNeed: 5,
				},
			},
		},
	}
	toppingMoreKeys := []string{"Recipe.Details.Ingredient"}
	invoice := invoicemodel.InvoiceCreate{
		Id: "Invoice001",
		InvoiceDetails: []invoicedetailmodel.InvoiceDetailCreate{
			{
				InvoiceId: "Invoice001",
				FoodId:    "Food001",
				SizeId:    "Size001",
				Toppings: &invoicedetailmodel.InvoiceDetailToppings{
					{
						Id: "Topping001",
					},
				},
				Amount: 1,
			},
		},
	}
	finalInvoice := invoicemodel.InvoiceCreate{
		Id:         "Invoice001",
		TotalPrice: 15000,
		InvoiceDetails: []invoicedetailmodel.InvoiceDetailCreate{
			{
				InvoiceId: "Invoice001",
				FoodId:    "Food001",
				SizeId:    "Size001",
				SizeName:  "Size Name",
				Toppings: &invoicedetailmodel.InvoiceDetailToppings{
					{
						Id:    "Topping001",
						Name:  "Topping Name",
						Price: 5000,
					},
				},
				Amount:    1,
				UnitPrice: 15000,
			},
		},
		MapIngredient: map[string]int{
			"Ingredient1": 15,
		},
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
			name: "Handle data failed because can not get size food",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoice,
			},
			mock: func() {
				sizeFoodStore.
					On("FindSizeFood",
						context.Background(),
						map[string]interface{}{
							"foodId": "Food001",
							"sizeId": "Size001",
						},
						sizeFoodMoreKeys).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle data failed because can not get topping",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoice,
			},
			mock: func() {
				sizeFoodStore.
					On("FindSizeFood",
						context.Background(),
						map[string]interface{}{
							"foodId": "Food001",
							"sizeId": "Size001",
						},
						sizeFoodMoreKeys).
					Return(&sizeFood, nil).
					Once()

				toppingStore.
					On("FindTopping",
						context.Background(),
						map[string]interface{}{
							"id": "Topping001",
						},
						toppingMoreKeys).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle data successfully",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoice,
			},
			mock: func() {
				sizeFoodStore.
					On("FindSizeFood",
						context.Background(),
						map[string]interface{}{
							"foodId": "Food001",
							"sizeId": "Size001",
						},
						sizeFoodMoreKeys).
					Return(&sizeFood, nil).
					Once()

				toppingStore.
					On("FindTopping",
						context.Background(),
						map[string]interface{}{
							"id": "Topping001",
						},
						toppingMoreKeys).
					Return(&topping, nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createInvoiceRepo{
				invoiceStore:       tt.fields.invoiceStore,
				invoiceDetailStore: tt.fields.invoiceDetailStore,
				customerStore:      tt.fields.customerStore,
				sizeFoodStore:      tt.fields.sizeFoodStore,
				foodStore:          tt.fields.foodStore,
				toppingStore:       tt.fields.toppingStore,
				ingredientStore:    tt.fields.ingredientStore,
				shopGeneralStore:   tt.fields.shopGeneralStore,
			}

			tt.mock()

			err := repo.HandleData(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "HandleData() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleData() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, *tt.args.data, finalInvoice, "Param = %v, want %v", *tt.args.data, finalInvoice)
			}
		})
	}
}

func Test_createInvoiceRepo_HandleIngredientTotalAmount(t *testing.T) {
	type fields struct {
		invoiceStore       InvoiceStore
		invoiceDetailStore InvoiceDetailStore
		customerStore      CustomerStore
		sizeFoodStore      SizeFoodStore
		foodStore          FoodStore
		toppingStore       ToppingStore
		ingredientStore    IngredientStore
		shopGeneralStore   ShopGeneralStore
	}
	type args struct {
		ctx                             context.Context
		invoiceId                       string
		ingredientTotalAmountNeedUpdate map[string]int
	}

	invoiceStore := new(mockInvoiceStore)
	invoiceDetailStore := new(mockInvoiceDetailStore)
	customerStore := new(mockCustomerStore)
	sizeFoodStore := new(mockSizeFoodStore)
	foodStore := new(mockFoodStore)
	toppingStore := new(mockToppingStore)
	ingredientStore := new(mockIngredientStore)
	shopGeneralStore := new(mockShopGeneralStore)

	ingredientMap := map[string]int{
		"Ingredient001": 12,
	}
	ingredient1 := ingredientmodel.Ingredient{
		Id:     "Ingredient001",
		Amount: 12,
	}
	outOfStockIngredient1 := ingredientmodel.Ingredient{
		Id:     "Ingredient001",
		Amount: 11,
	}
	ingredientUpdate1 := ingredientmodel.IngredientUpdateAmount{Amount: -12}
	var moreKeys []string
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Handle ingredient total amount failed because can not get ingredient",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:                             context.Background(),
				invoiceId:                       "Invoice001",
				ingredientTotalAmountNeedUpdate: ingredientMap,
			},
			mock: func() {
				ingredientStore.
					On("FindIngredient",
						context.Background(),
						map[string]interface{}{"id": "Ingredient001"},
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle ingredient total amount failed because ingredient is not enough",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:                             context.Background(),
				invoiceId:                       "Invoice001",
				ingredientTotalAmountNeedUpdate: ingredientMap,
			},
			mock: func() {
				ingredientStore.
					On("FindIngredient",
						context.Background(),
						map[string]interface{}{"id": "Ingredient001"},
						moreKeys).
					Return(&outOfStockIngredient1, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle ingredient total amount failed because can not update total amount ingredient",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:                             context.Background(),
				invoiceId:                       "Invoice001",
				ingredientTotalAmountNeedUpdate: ingredientMap,
			},
			mock: func() {
				ingredientStore.
					On("FindIngredient",
						context.Background(),
						map[string]interface{}{"id": "Ingredient001"},
						moreKeys).
					Return(&ingredient1, nil).
					Once()

				ingredientStore.
					On("UpdateAmountIngredient",
						context.Background(),
						"Ingredient001",
						&ingredientUpdate1).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle ingredient total amount successfully",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:                             context.Background(),
				invoiceId:                       "Invoice001",
				ingredientTotalAmountNeedUpdate: ingredientMap,
			},
			mock: func() {
				ingredientStore.
					On("FindIngredient",
						context.Background(),
						map[string]interface{}{"id": "Ingredient001"},
						moreKeys).
					Return(&ingredient1, nil).
					Once()

				ingredientStore.
					On("UpdateAmountIngredient",
						context.Background(),
						"Ingredient001",
						&ingredientUpdate1).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createInvoiceRepo{
				invoiceStore:       tt.fields.invoiceStore,
				invoiceDetailStore: tt.fields.invoiceDetailStore,
				customerStore:      tt.fields.customerStore,
				sizeFoodStore:      tt.fields.sizeFoodStore,
				foodStore:          tt.fields.foodStore,
				toppingStore:       tt.fields.toppingStore,
				ingredientStore:    tt.fields.ingredientStore,
				shopGeneralStore:   tt.fields.shopGeneralStore,
			}

			tt.mock()

			err := repo.HandleIngredientTotalAmount(tt.args.ctx, tt.args.invoiceId, tt.args.ingredientTotalAmountNeedUpdate)

			if tt.wantErr {
				assert.NotNil(t, err, "HandleInvoice() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleInvoice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createInvoiceRepo_HandleInvoice(t *testing.T) {
	type fields struct {
		invoiceStore       InvoiceStore
		invoiceDetailStore InvoiceDetailStore
		customerStore      CustomerStore
		sizeFoodStore      SizeFoodStore
		foodStore          FoodStore
		toppingStore       ToppingStore
		ingredientStore    IngredientStore
		shopGeneralStore   ShopGeneralStore
	}
	type args struct {
		ctx  context.Context
		data *invoicemodel.InvoiceCreate
	}

	invoiceStore := new(mockInvoiceStore)
	invoiceDetailStore := new(mockInvoiceDetailStore)
	customerStore := new(mockCustomerStore)
	sizeFoodStore := new(mockSizeFoodStore)
	foodStore := new(mockFoodStore)
	toppingStore := new(mockToppingStore)
	ingredientStore := new(mockIngredientStore)
	shopGeneralStore := new(mockShopGeneralStore)

	invoice := invoicemodel.InvoiceCreate{
		InvoiceDetails: []invoicedetailmodel.InvoiceDetailCreate{{}},
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
			name: "Create invoice failed because can not save invoice",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoice,
			},
			mock: func() {
				invoiceStore.
					On("CreateInvoice",
						context.Background(),
						&invoice).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create invoice failed because can not save invoice detail",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoice,
			},
			mock: func() {
				invoiceStore.
					On("CreateInvoice",
						context.Background(),
						&invoice).
					Return(nil).
					Once()

				invoiceDetailStore.
					On("CreateListImportNoteDetail",
						context.Background(),
						invoice.InvoiceDetails).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create invoice successfully",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoice,
			},
			mock: func() {
				invoiceStore.
					On("CreateInvoice",
						context.Background(),
						&invoice).
					Return(nil).
					Once()

				invoiceDetailStore.
					On("CreateListImportNoteDetail",
						context.Background(),
						invoice.InvoiceDetails).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createInvoiceRepo{
				invoiceStore:       tt.fields.invoiceStore,
				invoiceDetailStore: tt.fields.invoiceDetailStore,
				customerStore:      tt.fields.customerStore,
				sizeFoodStore:      tt.fields.sizeFoodStore,
				foodStore:          tt.fields.foodStore,
				toppingStore:       tt.fields.toppingStore,
				ingredientStore:    tt.fields.ingredientStore,
				shopGeneralStore:   tt.fields.shopGeneralStore,
			}

			tt.mock()

			err := repo.HandleInvoice(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "HandleInvoice() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleInvoice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createInvoiceRepo_UpdateCustomerPoint(t *testing.T) {
	type fields struct {
		invoiceStore       InvoiceStore
		invoiceDetailStore InvoiceDetailStore
		customerStore      CustomerStore
		sizeFoodStore      SizeFoodStore
		foodStore          FoodStore
		toppingStore       ToppingStore
		ingredientStore    IngredientStore
		shopGeneralStore   ShopGeneralStore
	}
	type args struct {
		ctx        context.Context
		customerId string
		data       customermodel.CustomerUpdatePoint
	}

	invoiceStore := new(mockInvoiceStore)
	invoiceDetailStore := new(mockInvoiceDetailStore)
	customerStore := new(mockCustomerStore)
	sizeFoodStore := new(mockSizeFoodStore)
	foodStore := new(mockFoodStore)
	toppingStore := new(mockToppingStore)
	ingredientStore := new(mockIngredientStore)
	shopGeneralStore := new(mockShopGeneralStore)

	amountUpdate := float32(12.3)
	customerUpdatePoint := customermodel.CustomerUpdatePoint{
		Amount: &amountUpdate,
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
			name: "Create invoice failed because can not save data to database",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:        context.Background(),
				customerId: "customer001",
				data:       customerUpdatePoint,
			},
			mock: func() {
				customerStore.
					On("UpdateCustomerPoint",
						context.Background(),
						"customer001",
						&customerUpdatePoint).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create invoice successfully",
			fields: fields{
				invoiceStore:       invoiceStore,
				invoiceDetailStore: invoiceDetailStore,
				customerStore:      customerStore,
				sizeFoodStore:      sizeFoodStore,
				foodStore:          foodStore,
				toppingStore:       toppingStore,
				ingredientStore:    ingredientStore,
				shopGeneralStore:   shopGeneralStore,
			},
			args: args{
				ctx:        context.Background(),
				customerId: "customer001",
				data:       customerUpdatePoint,
			},
			mock: func() {
				customerStore.
					On("UpdateCustomerPoint",
						context.Background(),
						"customer001",
						&customerUpdatePoint).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createInvoiceRepo{
				invoiceStore:       tt.fields.invoiceStore,
				invoiceDetailStore: tt.fields.invoiceDetailStore,
				customerStore:      tt.fields.customerStore,
				sizeFoodStore:      tt.fields.sizeFoodStore,
				foodStore:          tt.fields.foodStore,
				toppingStore:       tt.fields.toppingStore,
				ingredientStore:    tt.fields.ingredientStore,
				shopGeneralStore:   tt.fields.shopGeneralStore,
			}

			tt.mock()

			err := repo.UpdateCustomerPoint(tt.args.ctx, tt.args.customerId, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateCustomerPoint() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateCustomerPoint() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

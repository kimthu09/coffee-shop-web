package invoicerepo

import (
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/customerdebt/customerdebtmodel"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailmodel"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
	"context"
)

type InvoiceStore interface {
	CreateInvoice(
		ctx context.Context,
		data *invoicemodel.InvoiceCreate,
	) error
}

type InvoiceDetailStore interface {
	CreateListImportNoteDetail(
		ctx context.Context,
		data []invoicedetailmodel.InvoiceDetailCreate,
	) error
}

type CustomerStore interface {
	FindCustomer(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*customermodel.Customer, error)
	GetDebtCustomer(
		ctx context.Context,
		customerId string,
	) (*float32, error)
	UpdateCustomerDebt(
		ctx context.Context,
		id string,
		data *customermodel.CustomerUpdateDebt,
	) error
}

type CustomerDebtStore interface {
	CreateCustomerDebt(
		ctx context.Context,
		data *customerdebtmodel.CustomerDebtCreate,
	) error
}

type SizeFoodStore interface {
	FindSizeFood(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*sizefoodmodel.SizeFood, error)
}

type FoodStore interface {
	FindFood(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Food, error)
}

type ToppingStore interface {
	FindTopping(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Topping, error)
}

type createInvoiceRepo struct {
	invoiceStore       InvoiceStore
	invoiceDetailStore InvoiceDetailStore
	customerStore      CustomerStore
	customerDebtStore  CustomerDebtStore
	sizeFoodStore      SizeFoodStore
	foodStore          FoodStore
	toppingStore       ToppingStore
}

func NewCreateInvoiceRepo(
	invoiceStore InvoiceStore,
	invoiceDetailStore InvoiceDetailStore,
	customerStore CustomerStore,
	customerDebtStore CustomerDebtStore,
	sizeFoodStore SizeFoodStore,
	foodStore FoodStore,
	toppingStore ToppingStore) *createInvoiceRepo {
	return &createInvoiceRepo{
		invoiceStore:       invoiceStore,
		invoiceDetailStore: invoiceDetailStore,
		customerStore:      customerStore,
		customerDebtStore:  customerDebtStore,
		sizeFoodStore:      sizeFoodStore,
		foodStore:          foodStore,
		toppingStore:       toppingStore,
	}
}

func (repo *createInvoiceRepo) HandleCheckPermissionStatus(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	for _, invoiceDetail := range data.InvoiceDetails {
		if err := repo.checkPermissionStatusFood(ctx, invoiceDetail.FoodId); err != nil {
			return err
		}

		for _, topping := range *invoiceDetail.Toppings {
			if err := repo.checkPermissionStatusTopping(ctx, topping.Id); err != nil {
				return err
			}
		}
	}
	return nil
}

func (repo *createInvoiceRepo) checkPermissionStatusFood(
	ctx context.Context,
	foodId string) error {
	food, err := repo.foodStore.FindFood(
		ctx,
		map[string]interface{}{
			"Id": foodId,
		},
	)
	if err != nil {
		return err
	}

	if !food.IsActive {
		return invoicedetailmodel.ErrInvoiceDetailFoodIsInactive
	}
	return nil
}

func (repo *createInvoiceRepo) checkPermissionStatusTopping(
	ctx context.Context,
	toppingId string) error {
	topping, err := repo.toppingStore.FindTopping(
		ctx,
		map[string]interface{}{
			"Id": toppingId,
		},
	)
	if err != nil {
		return err
	}

	if !topping.IsActive {
		return invoicedetailmodel.ErrInvoiceDetailExistToppingIsInactive
	}
	return nil
}

func (repo *createInvoiceRepo) HandleData(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	totalPrice := float32(0)
	for i, invoiceDetail := range data.InvoiceDetails {
		sizeFood, errGetSizeFood := repo.sizeFoodStore.FindSizeFood(
			ctx,
			map[string]interface{}{
				"foodId": invoiceDetail.FoodId,
				"sizeId": invoiceDetail.SizeId,
			},
		)
		if errGetSizeFood != nil {
			return errGetSizeFood
		}
		data.InvoiceDetails[i].SizeName = sizeFood.Name

		priceToppings := float32(0)
		var toppings invoicedetailmodel.InvoiceDetailToppings
		for _, topping := range *invoiceDetail.Toppings {
			toppingModel, err := repo.toppingStore.FindTopping(
				ctx,
				map[string]interface{}{
					"id": topping.Id,
				},
			)
			if err != nil {
				return err
			}

			priceToppings += toppingModel.Price

			replaceTopping := invoicedetailmodel.InvoiceDetailTopping{
				Id:    topping.Id,
				Name:  toppingModel.Name,
				Price: toppingModel.Price,
			}
			toppings = append(toppings, replaceTopping)
		}
		*invoiceDetail.Toppings = toppings
		data.InvoiceDetails[i].UnitPrice = sizeFood.Price + priceToppings
		totalPrice += data.InvoiceDetails[i].UnitPrice * invoiceDetail.Amount
	}
	data.TotalPrice = totalPrice
	return nil
}

func (repo *createInvoiceRepo) getPriceFromTopping(
	ctx context.Context,
	toppingId string,
) (*float32, error) {
	topping, err := repo.toppingStore.FindTopping(
		ctx,
		map[string]interface{}{
			"id": toppingId,
		},
	)
	if err != nil {
		return nil, err
	}
	return &topping.Price, nil
}

func (repo *createInvoiceRepo) CreateCustomerDebt(
	ctx context.Context,
	supplierDebtId string,
	data *invoicemodel.InvoiceCreate) error {
	debtCurrent, err := repo.customerStore.GetDebtCustomer(
		ctx,
		*data.CustomerId)
	if err != nil {
		return err
	}

	amountBorrow := data.AmountDebt
	amountLeft := *debtCurrent + amountBorrow

	debtType := enum.Debt
	customerDebtCreate := customerdebtmodel.CustomerDebtCreate{
		Id:         supplierDebtId,
		CustomerId: *data.CustomerId,
		Amount:     amountBorrow,
		AmountLeft: amountLeft,
		DebtType:   &debtType,
		CreateBy:   data.CreateBy,
	}

	if err := repo.customerDebtStore.CreateCustomerDebt(
		ctx, &customerDebtCreate,
	); err != nil {
		return err
	}
	return nil
}

func (repo *createInvoiceRepo) UpdateDebtCustomer(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	customerUpdateDebt := customermodel.CustomerUpdateDebt{
		Amount: &data.AmountDebt,
	}
	if err := repo.customerStore.UpdateCustomerDebt(
		ctx, *data.CustomerId, &customerUpdateDebt,
	); err != nil {
		return err
	}
	return nil
}

func (repo *createInvoiceRepo) HandleInvoice(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	if err := repo.createInvoice(ctx, data); err != nil {
		return err
	}

	if err := repo.createInvoiceDetails(ctx, data.InvoiceDetails); err != nil {
		return err
	}
	return nil
}

func (repo *createInvoiceRepo) createInvoice(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	if err := repo.invoiceStore.CreateInvoice(ctx, data); err != nil {
		return err
	}
	return nil
}

func (repo *createInvoiceRepo) createInvoiceDetails(
	ctx context.Context,
	data []invoicedetailmodel.InvoiceDetailCreate) error {
	if err := repo.invoiceDetailStore.CreateListImportNoteDetail(
		ctx, data,
	); err != nil {
		return err
	}
	return nil
}

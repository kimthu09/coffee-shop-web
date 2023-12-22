package invoicebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"coffee_shop_management_backend/module/shopgeneral/shopgeneralmodel"
	"context"
)

type CreateInvoiceRepo interface {
	GetShopGeneral(
		ctx context.Context,
	) (*shopgeneralmodel.ShopGeneral, error)
	HandleCheckPermissionStatus(
		ctx context.Context,
		data *invoicemodel.InvoiceCreate,
	) error
	HandleData(
		ctx context.Context,
		data *invoicemodel.InvoiceCreate,
	) error
	FindCustomer(
		ctx context.Context,
		customerId string,
	) (*customermodel.Customer, error)
	UpdateCustomerPoint(
		ctx context.Context,
		customerId string,
		data customermodel.CustomerUpdatePoint,
	) error
	HandleInvoice(
		ctx context.Context,
		data *invoicemodel.InvoiceCreate,
	) error
	HandleIngredientTotalAmount(
		ctx context.Context,
		invoiceId string,
		ingredientTotalAmountNeedUpdate map[string]int,
	) error
}

type createInvoiceBiz struct {
	gen       generator.IdGenerator
	repo      CreateInvoiceRepo
	requester middleware.Requester
}

func NewCreateInvoiceBiz(
	gen generator.IdGenerator,
	repo CreateInvoiceRepo,
	requester middleware.Requester) *createInvoiceBiz {
	return &createInvoiceBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createInvoiceBiz) CreateInvoice(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	if !biz.requester.IsHasFeature(common.InvoiceCreateFeatureCode) {
		return invoicemodel.ErrInvoiceCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleInvoiceId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.HandleCheckPermissionStatus(ctx, data); err != nil {
		return err
	}

	if err := biz.repo.HandleData(ctx, data); err != nil {
		return err
	}

	general, errGetShopGeneral := biz.repo.GetShopGeneral(ctx)
	if errGetShopGeneral != nil {
		return errGetShopGeneral
	}
	data.ShopName = general.Name
	data.ShopPhone = general.Phone
	data.ShopAddress = general.Address
	data.ShopPassWifi = general.WifiPass
	if err := biz.repo.HandleIngredientTotalAmount(
		ctx, data.Id, data.MapIngredient); err != nil {
		return err
	}

	if data.CustomerId != nil {
		customer, errGetCustomer := biz.repo.FindCustomer(ctx, *data.CustomerId)
		if errGetCustomer != nil {
			return errGetCustomer
		}

		data.Customer.Id = customer.Id
		data.Customer.Name = customer.Name
		data.Customer.Phone = customer.Phone

		priceUseForPoint := float32(0)
		pointUse := float32(0)
		if data.IsUsePoint {
			if float32(data.TotalPrice) >= customer.Point*general.UsePointPercent {
				pointUse = customer.Point
				priceUseForPoint = customer.Point * general.UsePointPercent
			} else {
				pointUse = float32(data.TotalPrice) / general.UsePointPercent
				priceUseForPoint = float32(data.TotalPrice)
			}
		}
		priceUseForPointInt := common.RoundToInt(priceUseForPoint)

		amountPointNeedUpdate :=
			float32(data.AmountReceived)*general.AccumulatePointPercent - pointUse
		common.CustomRound(&amountPointNeedUpdate)
		customerUpdatePoint := customermodel.CustomerUpdatePoint{
			Amount: &amountPointNeedUpdate,
		}
		if err := biz.repo.UpdateCustomerPoint(
			ctx, *data.CustomerId, customerUpdatePoint); err != nil {
			return err
		}

		common.CustomRound(&priceUseForPoint)
		data.AmountReceived = data.TotalPrice - priceUseForPointInt
		data.AmountPriceUsePoint = priceUseForPointInt

	} else {
		if data.IsUsePoint {
			return invoicemodel.ErrInvoiceNotHaveCustomerToUsePoint
		}
	}

	if err := biz.repo.HandleInvoice(ctx, data); err != nil {
		return err
	}
	return nil
}

func handleInvoiceId(gen generator.IdGenerator, data *invoicemodel.InvoiceCreate) error {
	idInvoice, err := gen.GenerateId()
	if err != nil {
		return err
	}
	data.Id = idInvoice

	for i := range data.InvoiceDetails {
		data.InvoiceDetails[i].InvoiceId = idInvoice
	}

	return err
}

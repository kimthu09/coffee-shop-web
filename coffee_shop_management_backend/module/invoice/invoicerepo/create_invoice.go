package invoicerepo

import (
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailmodel"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/shopgeneral/shopgeneralmodel"
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
	UpdateCustomerPoint(
		ctx context.Context,
		id string,
		data *customermodel.CustomerUpdatePoint,
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

type IngredientStore interface {
	FindIngredient(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*ingredientmodel.Ingredient, error)
	UpdateAmountIngredient(
		ctx context.Context,
		id string,
		data *ingredientmodel.IngredientUpdateAmount,
	) error
}

type ShopGeneralStore interface {
	FindShopGeneral(
		ctx context.Context,
	) (*shopgeneralmodel.ShopGeneral, error)
}

type createInvoiceRepo struct {
	invoiceStore       InvoiceStore
	invoiceDetailStore InvoiceDetailStore
	customerStore      CustomerStore
	sizeFoodStore      SizeFoodStore
	foodStore          FoodStore
	toppingStore       ToppingStore
	ingredientStore    IngredientStore
	shopGeneralStore   ShopGeneralStore
}

func NewCreateInvoiceRepo(
	invoiceStore InvoiceStore,
	invoiceDetailStore InvoiceDetailStore,
	customerStore CustomerStore,
	sizeFoodStore SizeFoodStore,
	foodStore FoodStore,
	toppingStore ToppingStore,
	ingredientStore IngredientStore,
	shopGeneralStore ShopGeneralStore) *createInvoiceRepo {
	return &createInvoiceRepo{
		invoiceStore:       invoiceStore,
		invoiceDetailStore: invoiceDetailStore,
		customerStore:      customerStore,
		sizeFoodStore:      sizeFoodStore,
		foodStore:          foodStore,
		toppingStore:       toppingStore,
		ingredientStore:    ingredientStore,
		shopGeneralStore:   shopGeneralStore,
	}
}

func (repo *createInvoiceRepo) GetShopGeneral(
	ctx context.Context) (*shopgeneralmodel.ShopGeneral, error) {
	if data, err := repo.shopGeneralStore.FindShopGeneral(ctx); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func (repo *createInvoiceRepo) HandleCheckPermissionStatus(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	for i, invoiceDetail := range data.InvoiceDetails {
		if food, err := repo.getFood(ctx, invoiceDetail.FoodId); err != nil {
			return err
		} else {
			data.InvoiceDetails[i].FoodName = food.Name
		}

		for _, topping := range *invoiceDetail.Toppings {
			if err := repo.checkPermissionStatusTopping(ctx, topping.Id); err != nil {
				return err
			}
		}
	}
	return nil
}

func (repo *createInvoiceRepo) getFood(
	ctx context.Context,
	foodId string) (*productmodel.Food, error) {
	food, err := repo.foodStore.FindFood(
		ctx,
		map[string]interface{}{
			"Id": foodId,
		},
	)
	if err != nil {
		return nil, err
	}

	if !food.IsActive {
		return nil, invoicedetailmodel.ErrInvoiceDetailFoodIsInactive
	}

	return food, nil
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
	totalPrice := 0

	mapTopping := make(map[string]int)
	mapToppingName := make(map[string]string)
	mapToppingPrice := make(map[string]int)
	mapFood := make(map[string]map[string]int)
	mapFoodSizeName := make(map[string]map[string]string)
	mapFoodSizePrice := make(map[string]map[string]int)

	for _, detail := range data.InvoiceDetails {
		for _, topping := range *detail.Toppings {
			mapTopping[topping.Id]++
		}
		if mapFood[detail.FoodId] == nil {
			mapFood[detail.FoodId] = make(map[string]int)
		}
		mapFood[detail.FoodId][detail.SizeId]++
	}

	mapIngredient := make(map[string]int)

	for keyFood, mapSize := range mapFood {
		for keySize, value := range mapSize {
			sizeFood, errGetSizeFood := repo.sizeFoodStore.FindSizeFood(
				ctx,
				map[string]interface{}{
					"foodId": keyFood,
					"sizeId": keySize,
				},
				"Recipe.Details.Ingredient",
			)
			if errGetSizeFood != nil {
				return errGetSizeFood
			}

			if mapFoodSizeName[keyFood] == nil {
				mapFoodSizeName[keyFood] = make(map[string]string)
			}
			mapFoodSizeName[keyFood][keySize] = sizeFood.Name

			if mapFoodSizePrice[keyFood] == nil {
				mapFoodSizePrice[keyFood] = make(map[string]int)
			}
			mapFoodSizePrice[keyFood][keySize] += sizeFood.Cost

			for _, recipeDetail := range sizeFood.Recipe.Details {
				mapIngredient[recipeDetail.IngredientId] +=
					recipeDetail.AmountNeed * value
			}
		}
	}

	for key, value := range mapTopping {
		topping, errGetTopping := repo.toppingStore.FindTopping(
			ctx,
			map[string]interface{}{
				"id": key,
			},
			"Recipe.Details.Ingredient",
		)
		if errGetTopping != nil {
			return errGetTopping
		}

		mapToppingName[key] = topping.Name
		mapToppingPrice[key] = topping.Cost

		for _, recipeDetail := range topping.Recipe.Details {
			mapIngredient[recipeDetail.IngredientId] += recipeDetail.AmountNeed * value
		}
	}

	data.MapIngredient = mapIngredient
	for i, invoiceDetail := range data.InvoiceDetails {
		data.InvoiceDetails[i].SizeName = mapFoodSizeName[invoiceDetail.FoodId][invoiceDetail.SizeId]

		priceToppings := 0
		var toppings invoicedetailmodel.InvoiceDetailToppings

		for _, topping := range *invoiceDetail.Toppings {
			priceToppings += mapToppingPrice[topping.Id]

			simpleTopping := invoicedetailmodel.InvoiceDetailTopping{
				Id:    topping.Id,
				Name:  mapToppingName[topping.Id],
				Price: mapToppingPrice[topping.Id],
			}
			toppings = append(toppings, simpleTopping)
		}

		*invoiceDetail.Toppings = toppings
		data.InvoiceDetails[i].UnitPrice =
			mapFoodSizePrice[invoiceDetail.FoodId][invoiceDetail.SizeId] + priceToppings
		totalPrice += data.InvoiceDetails[i].UnitPrice * invoiceDetail.Amount
	}

	data.TotalPrice = totalPrice

	return nil
}

func (repo *createInvoiceRepo) FindCustomer(
	ctx context.Context,
	customerId string) (*customermodel.Customer, error) {
	customer, err := repo.customerStore.FindCustomer(
		ctx, map[string]interface{}{"id": customerId},
	)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (repo *createInvoiceRepo) UpdateCustomerPoint(
	ctx context.Context,
	customerId string,
	data customermodel.CustomerUpdatePoint) error {
	if err := repo.customerStore.UpdateCustomerPoint(
		ctx, customerId, &data,
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

func (repo *createInvoiceRepo) HandleIngredientTotalAmount(
	ctx context.Context,
	invoiceId string,
	ingredientTotalAmountNeedUpdate map[string]int) error {
	for key, value := range ingredientTotalAmountNeedUpdate {
		ingredient, errGetIngredient := repo.ingredientStore.FindIngredient(
			ctx, map[string]interface{}{"id": key})
		if errGetIngredient != nil {
			return errGetIngredient
		}

		amountLeft := ingredient.Amount - value
		if amountLeft < 0 {
			return invoicemodel.ErrInvoiceIngredientIsNotEnough
		}

		ingredientUpdate := ingredientmodel.IngredientUpdateAmount{Amount: -value}
		if err := repo.ingredientStore.UpdateAmountIngredient(
			ctx, key, &ingredientUpdate,
		); err != nil {
			return err
		}
	}
	return nil
}

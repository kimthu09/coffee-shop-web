package productrepo

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type SeeDetailToppingStore interface {
	FindTopping(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Topping, error)
}

type seeDetailToppingRepo struct {
	store SeeDetailToppingStore
}

func NewSeeDetailToppingRepo(store SeeDetailToppingStore) *seeDetailToppingRepo {
	return &seeDetailToppingRepo{store: store}
}

func (repo *seeDetailToppingRepo) SeeDetailTopping(
	ctx context.Context,
	toppingId string) (*productmodel.Topping, error) {
	result, err := repo.store.FindTopping(
		ctx,
		map[string]interface{}{"id": toppingId},
		"Recipe.Details.Ingredient")

	if err != nil {
		return nil, err
	}

	return result, nil
}

package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"context"
)

type UpdateToppingRepo interface {
	FindTopping(
		ctx context.Context,
		id string,
	) (*productmodel.Topping, error)
	UpdateTopping(
		ctx context.Context,
		id string,
		data *productmodel.ToppingUpdateInfo,
	) error
	UpdateRecipeDetailsOfRecipe(
		ctx context.Context,
		recipeId string,
		deletedRecipeDetails []recipedetailmodel.RecipeDetail,
		updatedRecipeDetails []recipedetailmodel.RecipeDetailUpdate,
		createdRecipeDetails []recipedetailmodel.RecipeDetailCreate,
	) error
	FindRecipeDetails(
		ctx context.Context,
		recipeId string,
	) ([]recipedetailmodel.RecipeDetail, error)
}

type updateToppingBiz struct {
	repo      UpdateToppingRepo
	requester middleware.Requester
}

func NewUpdateToppingBiz(
	repo UpdateToppingRepo,
	requester middleware.Requester) *updateToppingBiz {
	return &updateToppingBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *updateToppingBiz) UpdateTopping(
	ctx context.Context,
	id string,
	data *productmodel.ToppingUpdateInfo) error {
	if !biz.requester.IsHasFeature(common.ToppingUpdateInfoFeatureCode) {
		return productmodel.ErrToppingUpdateInfoNoPermission
	}

	//validate data
	if err := data.Validate(); err != nil {
		return err
	}

	//Check status
	currentTopping, err := biz.repo.FindTopping(ctx, id)
	if err != nil {
		return err
	}

	if !currentTopping.IsActive {
		return common.ErrNoPermission(productmodel.ErrProductInactive)
	}

	//Store data
	if err := biz.repo.UpdateTopping(ctx, id, data); err != nil {
		return err
	}

	if data.Recipe != nil {
		if err := biz.updateRecipe(ctx, currentTopping.RecipeId, data); err != nil {
			return err
		}
	}
	return nil
}

func (biz *updateToppingBiz) updateRecipe(
	ctx context.Context,
	recipeId string,
	data *productmodel.ToppingUpdateInfo) error {
	currentRecipeDetails, err := biz.repo.FindRecipeDetails(ctx, recipeId)
	if err != nil {
		return err
	}

	deletedRecipeDetails, updatedRecipeDetails, createdRecipeDetails := classifyDetails(
		recipeId,
		currentRecipeDetails,
		data.Recipe.Details,
	)

	if err := biz.repo.UpdateRecipeDetailsOfRecipe(
		ctx,
		recipeId,
		deletedRecipeDetails,
		updatedRecipeDetails,
		createdRecipeDetails,
	); err != nil {
		return err
	}
	return nil
}

func classifyDetails(
	recipeId string,
	currents []recipedetailmodel.RecipeDetail,
	updates []recipedetailmodel.RecipeDetailUpdate) (
	[]recipedetailmodel.RecipeDetail,
	[]recipedetailmodel.RecipeDetailUpdate,
	[]recipedetailmodel.RecipeDetailCreate) {
	var deletedRecipeDetails []recipedetailmodel.RecipeDetail
	var updatedRecipeDetails []recipedetailmodel.RecipeDetailUpdate
	var createdRecipeDetails []recipedetailmodel.RecipeDetailCreate

	mapExistRecipeDetail := make(map[string]int)
	mapRecipeDetailUpdate := make(map[string]recipedetailmodel.RecipeDetailUpdate)
	mapRecipeDetail := make(map[string]recipedetailmodel.RecipeDetail)
	for _, v := range updates {
		mapExistRecipeDetail[v.IngredientId]++
		mapRecipeDetailUpdate[v.IngredientId] = v
	}
	for _, v := range currents {
		mapExistRecipeDetail[v.IngredientId]--
		mapRecipeDetail[v.IngredientId] = v
	}
	for key, value := range mapExistRecipeDetail {
		if value == -1 {
			deletedRecipeDetails = append(
				deletedRecipeDetails,
				mapRecipeDetail[key],
			)
		}
		if value == 0 {
			updatedRecipeDetails = append(
				updatedRecipeDetails,
				mapRecipeDetailUpdate[key],
			)
		}
		if value == 1 {
			recipeDetailNeedCreate := mapRecipeDetailUpdate[key]
			recipeDetailCreate := recipedetailmodel.RecipeDetailCreate{
				RecipeId:     recipeId,
				IngredientId: recipeDetailNeedCreate.IngredientId,
				AmountNeed:   recipeDetailNeedCreate.AmountNeed,
			}
			createdRecipeDetails = append(
				createdRecipeDetails,
				recipeDetailCreate,
			)
		}
	}

	return deletedRecipeDetails, updatedRecipeDetails, createdRecipeDetails
}

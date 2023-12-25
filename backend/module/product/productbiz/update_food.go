package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/category/categorymodel"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
	"context"
)

type UpdateFoodRepo interface {
	FindFood(
		ctx context.Context,
		id string,
	) (*productmodel.Food, error)
	FindCategories(
		ctx context.Context,
		foodId string,
	) ([]categorymodel.SimpleCategoryWithId, error)
	HandleCategory(
		ctx context.Context,
		foodId string,
		deletedCategoryFood []categorymodel.SimpleCategoryWithId,
		createdCategoryFood []categorymodel.SimpleCategoryWithId,
	) error
	FindSizeFoods(
		ctx context.Context,
		foodId string,
	) ([]sizefoodmodel.SizeFood, error)
	FindRecipeDetails(
		ctx context.Context,
		recipeId string,
	) ([]recipedetailmodel.RecipeDetail, error)
	HandleSizeFoods(
		ctx context.Context,
		foodId string,
		deletedSizeFood []sizefoodmodel.SizeFood,
		updatedSizeFood []sizefoodmodel.SizeFoodUpdate,
		mapDeletedRecipeDetails map[string][]recipedetailmodel.RecipeDetail,
		mapUpdatedRecipeDetails map[string][]recipedetailmodel.RecipeDetailUpdate,
		mapCreatedRecipeDetails map[string][]recipedetailmodel.RecipeDetailCreate,
		createdSizeFood []sizefoodmodel.SizeFoodCreate) error
	UpdateFood(
		ctx context.Context,
		id string,
		data *productmodel.FoodUpdateInfo,
	) error
}

type updateFoodBiz struct {
	gen       generator.IdGenerator
	repo      UpdateFoodRepo
	requester middleware.Requester
}

func NewUpdateFoodBiz(
	gen generator.IdGenerator,
	repo UpdateFoodRepo,
	requester middleware.Requester) *updateFoodBiz {
	return &updateFoodBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *updateFoodBiz) UpdateFood(
	ctx context.Context,
	id string,
	data *productmodel.FoodUpdateInfo) error {
	if !biz.requester.IsHasFeature(common.FoodUpdateInfoFeatureCode) {
		return productmodel.ErrFoodUpdateInfoNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	result, err := biz.repo.FindFood(ctx, id)
	if err != nil {
		return err
	}

	if !result.IsActive {
		return common.ErrNoPermission(productmodel.ErrProductInactive)
	}

	//handle update category
	if data.Categories != nil {
		var deletedCategories []categorymodel.SimpleCategoryWithId
		var createdCategories []categorymodel.SimpleCategoryWithId

		///handle get change of amount product
		simpleCategories, err := biz.repo.FindCategories(ctx, id)
		if err != nil {
			return err
		}

		mapCategoriesAmountProduct := make(map[string]int)
		for _, v := range simpleCategories {
			mapCategoriesAmountProduct[v.CategoryId]--
		}
		for _, v := range *data.Categories {
			mapCategoriesAmountProduct[v]++
		}

		for key, value := range mapCategoriesAmountProduct {
			if value == 0 {
				continue
			} else {
				simpleCategory := categorymodel.SimpleCategoryWithId{
					CategoryId: key,
				}
				if value == 1 {
					createdCategories = append(createdCategories, simpleCategory)
				} else if value == -1 {
					deletedCategories = append(deletedCategories, simpleCategory)
				}
			}
		}

		if err := biz.repo.HandleCategory(
			ctx,
			id,
			deletedCategories,
			createdCategories); err != nil {
			return err
		}
	}

	if data.Sizes != nil {
		var deletedSizes []sizefoodmodel.SizeFood
		var updatedSizes []sizefoodmodel.SizeFoodUpdate
		var createdSizes []sizefoodmodel.SizeFoodCreate
		mapDeletedRecipeDetails := make(map[string][]recipedetailmodel.RecipeDetail)
		mapUpdatedRecipeDetails := make(map[string][]recipedetailmodel.RecipeDetailUpdate)
		mapCreatedRecipeDetails := make(map[string][]recipedetailmodel.RecipeDetailCreate)

		///get current size foods
		currentSizeFoods, err := biz.repo.FindSizeFoods(ctx, id)
		if err != nil {
			return err
		}

		///classify size food to handle
		mapExistSize := make(map[string]int)
		mapSizeFoodUpdate := make(map[string]sizefoodmodel.SizeFoodUpdate)
		mapSizeFood := make(map[string]sizefoodmodel.SizeFood)

		for _, v := range currentSizeFoods {
			mapExistSize[v.SizeId]--
			mapSizeFood[v.SizeId] = v
		}
		for _, v := range *data.Sizes {
			if v.SizeId == nil {
				sizeCreate, errCreate := getSizeFoodCreateFromSizeFoodUpdate(biz.gen, id, v)
				if errCreate != nil {
					return errCreate
				}

				if err := sizeCreate.Validate(); err != nil {
					return err
				}

				createdSizes = append(createdSizes, *sizeCreate)
			} else {
				mapExistSize[*v.SizeId]++
				if mapExistSize[*v.SizeId] == 1 {
					return common.ErrInternal(productmodel.ErrFoodSizeIdInvalid)
				}
				mapSizeFoodUpdate[*v.SizeId] = v
			}
		}

		for key, value := range mapExistSize {
			currentSize := mapSizeFood[key]
			if value == 0 {
				size := mapSizeFoodUpdate[key]
				size.RecipeId = &currentSize.RecipeId
				updatedSizes = append(updatedSizes, size)

				if size.Recipe != nil &&
					size.Recipe.Details != nil &&
					len(size.Recipe.Details) != 0 {
					currentReceiptDetails, err := biz.repo.FindRecipeDetails(
						ctx,
						currentSize.RecipeId)
					if err != nil {
						return err
					}

					deletedRecipeDetails, updatedRecipeDetails, createRecipeDetails :=
						classifyDetails(
							*size.RecipeId,
							currentReceiptDetails,
							size.Recipe.Details,
						)

					mapDeletedRecipeDetails[*size.RecipeId] = deletedRecipeDetails
					mapUpdatedRecipeDetails[*size.RecipeId] = updatedRecipeDetails
					mapCreatedRecipeDetails[*size.RecipeId] = createRecipeDetails
				}
			}
			if value == -1 {
				deletedSizes = append(deletedSizes, currentSize)
			}
		}

		if err := biz.repo.HandleSizeFoods(
			ctx,
			id,
			deletedSizes,
			updatedSizes,
			mapDeletedRecipeDetails,
			mapUpdatedRecipeDetails,
			mapCreatedRecipeDetails,
			createdSizes); err != nil {
			return err
		}
	}

	if err := biz.repo.UpdateFood(ctx, id, data); err != nil {
		return err
	}

	return nil
}

func getSizeFoodCreateFromSizeFoodUpdate(
	gen generator.IdGenerator,
	foodId string,
	size sizefoodmodel.SizeFoodUpdate) (*sizefoodmodel.SizeFoodCreate, error) {
	sizeId, err := gen.GenerateId()
	if err != nil {
		return nil, err
	}

	name := ""
	if size.Name != nil {
		name = *size.Name
	}

	cost := -1
	if size.Cost != nil {
		cost = *size.Cost
	}

	price := -1
	if size.Price != nil {
		price = *size.Price
	}

	var recipe *recipemodel.RecipeCreate
	if size.Recipe != nil && size.Recipe.Details != nil {
		recipeId, err := gen.GenerateId()
		if err != nil {
			return nil, err
		}
		var details []recipedetailmodel.RecipeDetailCreate
		for _, detailUpdate := range size.Recipe.Details {
			ingredientId := detailUpdate.IngredientId
			amountNeed := detailUpdate.AmountNeed
			detail := recipedetailmodel.RecipeDetailCreate{
				RecipeId:     recipeId,
				IngredientId: ingredientId,
				AmountNeed:   amountNeed,
			}
			details = append(details, detail)
		}
		recipe = &recipemodel.RecipeCreate{
			Id:      recipeId,
			Details: details,
		}
	}

	receiptId := ""
	if recipe != nil {
		receiptId = recipe.Id
	}

	sizeCreate := sizefoodmodel.SizeFoodCreate{
		FoodId:   foodId,
		SizeId:   sizeId,
		Name:     name,
		Cost:     cost,
		Price:    price,
		RecipeId: receiptId,
		Recipe:   recipe,
	}

	return &sizeCreate, nil
}

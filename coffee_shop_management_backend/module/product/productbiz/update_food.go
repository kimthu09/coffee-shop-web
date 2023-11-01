package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/asyncjob"
	"coffee_shop_management_backend/module/category/categorymodel"
	"coffee_shop_management_backend/module/categoryfood/categoryfoodmodel"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
	"context"
)

type UpdateFoodStore interface {
	FindFood(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Food, error)
	UpdateFood(
		ctx context.Context,
		id string,
		data *productmodel.FoodUpdate,
	) error
}

type CreateOrDeleteCategoryFoodStore interface {
	FindListCategories(
		ctx context.Context,
		foodId string,
	) (*[]categorymodel.SimpleCategory, error)
	CreateCategoryFood(
		ctx context.Context,
		data *categoryfoodmodel.CategoryFoodCreate,
	) error
	DeleteCategoryFood(
		ctx context.Context,
		conditions map[string]interface{},
	) error
}

type UpdateSizeFoodStore interface {
	FindListSizeFood(
		ctx context.Context,
		foodId string,
	) (*[]sizefoodmodel.SizeFood, error)
	CreateSizeFood(
		ctx context.Context,
		data *sizefoodmodel.SizeFoodCreate,
	) error
	DeleteSizeFood(
		ctx context.Context,
		conditions map[string]interface{},
	) error
	UpdateSizeFood(
		ctx context.Context,
		foodId string,
		sizeId string,
		data *sizefoodmodel.SizeFoodUpdate,
	) error
}

type UpdateRecipeStore interface {
	CreateRecipe(
		ctx context.Context,
		data *recipemodel.RecipeCreate,
	) error
	DeleteRecipe(
		ctx context.Context,
		conditions map[string]interface{},
	) error
}

type UpdateRecipeDetailStore interface {
	UpdateRecipeDetail(
		ctx context.Context,
		idRecipe string,
		idIngredient string,
		data *recipedetailmodel.RecipeDetailUpdate,
	) error
	CreateListRecipeDetail(
		ctx context.Context,
		data []recipedetailmodel.RecipeDetailCreate,
	) error
	FindListRecipeDetail(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*[]recipedetailmodel.RecipeDetail, error)
	DeleteRecipeDetail(
		ctx context.Context,
		conditions map[string]interface{},
	) error
}

type updateFoodBiz struct {
	foodStore         UpdateFoodStore
	categoryFoodStore CreateOrDeleteCategoryFoodStore
	categoryStore     UpdateCategoryStore
	sizeFoodStore     UpdateSizeFoodStore
	recipeStore       UpdateRecipeStore
	ingredientStore   CheckIngredientStore
	recipeDetailStore UpdateRecipeDetailStore
}

func NewUpdateFoodBiz(
	foodStore UpdateFoodStore,
	categoryFoodStore CreateOrDeleteCategoryFoodStore,
	categoryStore UpdateCategoryStore,
	sizeFoodStore UpdateSizeFoodStore,
	recipeStore UpdateRecipeStore,
	ingredientStore CheckIngredientStore,
	recipeDetailStore UpdateRecipeDetailStore) *updateFoodBiz {
	return &updateFoodBiz{
		foodStore:         foodStore,
		categoryFoodStore: categoryFoodStore,
		categoryStore:     categoryStore,
		sizeFoodStore:     sizeFoodStore,
		recipeStore:       recipeStore,
		ingredientStore:   ingredientStore,
		recipeDetailStore: recipeDetailStore,
	}
}

func (biz *updateFoodBiz) UpdateFood(
	ctx context.Context,
	id string,
	data *productmodel.FoodUpdate) error {
	//validate data
	if err := data.Validate(); err != nil {
		return err
	}

	result, err := biz.foodStore.FindFood(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if !result.IsActive {
		return common.ErrNoPermission(productmodel.ErrProductInactive)
	}

	//handle update category
	var categoryProductJobs []asyncjob.Job
	var categoryJobs []asyncjob.Job
	if data.Categories != nil {
		///check category exists
		for _, v := range *data.Categories {
			if _, err := biz.categoryStore.FindCategory(
				ctx,
				map[string]interface{}{"id": v}); err != nil {
				return err
			}
		}

		///handle get change of amount product
		var simpleCategories *[]categorymodel.SimpleCategory
		simpleCategories, err = biz.categoryFoodStore.FindListCategories(ctx, id)
		if err != nil {
			return err
		}

		mapCategoriesAmountProduct := make(map[string]int)
		for _, v := range *simpleCategories {
			mapCategoriesAmountProduct[v.CategoryId]--
		}
		for _, v := range *data.Categories {
			mapCategoriesAmountProduct[v]++
		}

		///handle define job create or delete category product
		for key, value := range mapCategoriesAmountProduct {
			tempJob := asyncjob.NewJob(func(
				productId string,
				categoryId string,
				amount int,
			) func(ctx context.Context) error {
				return func(ctx context.Context) error {
					if amount == 0 {
						return nil
					}
					if amount == 1 {
						categoryFoodCreate := categoryfoodmodel.CategoryFoodCreate{
							FoodId:     &productId,
							CategoryId: &categoryId,
						}
						return biz.categoryFoodStore.CreateCategoryFood(
							ctx,
							&categoryFoodCreate)
					}
					if amount == -1 {
						return biz.categoryFoodStore.DeleteCategoryFood(
							ctx,
							map[string]interface{}{
								"productId":  productId,
								"categoryId": categoryId,
							})
					}
					return nil
				}
			}(id, key, value))
			categoryProductJobs = append(categoryProductJobs, tempJob)
		}

		///handle define job update amount product of category
		for key, value := range mapCategoriesAmountProduct {
			tempJob := asyncjob.NewJob(func(
				categoryId string,
				amount int,
			) func(ctx context.Context) error {
				return func(ctx context.Context) error {
					if value == 0 {
						return nil
					}
					categoryUpdateAmountProduct := categorymodel.CategoryUpdateAmountProduct{
						AmountProduct: &amount,
					}
					return biz.categoryStore.UpdateAmountProductCategory(
						ctx,
						categoryId,
						&categoryUpdateAmountProduct)
				}
			}(key, value))
			categoryJobs = append(categoryJobs, tempJob)
		}
	}

	//handle update size food
	var sizeFoodJobs []asyncjob.Job
	var recipeDetailJobs []asyncjob.Job
	var recipeJobs []asyncjob.Job
	if data.Sizes != nil {
		///check ingredients exists
		for _, size := range *data.Sizes {
			if *size.Recipe.Details == nil {
				continue
			}
			for _, recipeDetail := range *size.Recipe.Details {
				if _, err := biz.ingredientStore.FindIngredient(
					ctx,
					map[string]interface{}{"id": recipeDetail.IngredientId},
				); err != nil {
					return err
				}
			}
		}

		///get current size foods
		currentSizeFoods, err := biz.sizeFoodStore.FindListSizeFood(
			ctx,
			id,
		)
		if err != nil {
			return err
		}

		///classify size food to handle
		mapExistSize := make(map[string]int)
		mapSizeFoodUpdate := make(map[string]sizefoodmodel.SizeFoodUpdate)
		mapSizeFood := make(map[string]sizefoodmodel.SizeFood)
		var sizeFoodNeedCreate []sizefoodmodel.SizeFoodUpdate

		for _, v := range *currentSizeFoods {
			mapExistSize[v.SizeId]--
			mapSizeFood[v.SizeId] = v
		}

		for _, v := range *data.Sizes {
			if v.SizeId == nil {
				sizeFoodNeedCreate = append(sizeFoodNeedCreate, v)
			} else {
				mapExistSize[*v.SizeId]++
				if mapExistSize[*v.SizeId] == 1 {
					return common.ErrInternal(productmodel.ErrSizeIdInvalid)
				}
				mapSizeFoodUpdate[*v.SizeId] = v
			}
		}

		///create size food
		for _, size := range sizeFoodNeedCreate {
			sizeId, err := common.GenerateId()
			if err != nil {
				return err
			}

			name := ""
			if size.Name != nil {
				name = *size.Name
			}

			cost := -1.0
			if size.Cost != nil {
				cost = *size.Cost
			}

			price := -1.0
			if size.Price != nil {
				cost = *size.Price
			}

			var recipe *recipemodel.RecipeCreate
			if size.Recipe != nil && size.Recipe.Details != nil {
				recipeId, err := common.GenerateId()
				if err != nil {
					return err
				}
				var details []recipedetailmodel.RecipeDetailCreate
				for _, detailUpdate := range *size.Recipe.Details {
					ingredientId := ""
					amountNeed := float32(-1.0)
					if detailUpdate.IngredientId != nil {
						ingredientId = *detailUpdate.IngredientId
					}
					if detailUpdate.AmountNeed != nil {
						amountNeed = *detailUpdate.AmountNeed
					}
					detail := recipedetailmodel.RecipeDetailCreate{
						RecipeId:     recipeId,
						IngredientId: ingredientId,
						AmountNeed:   amountNeed,
					}
					details = append(details, detail)
				}
				recipe = &recipemodel.RecipeCreate{
					Id:      recipeId,
					Details: &details,
				}
			}

			receiptId := ""
			if recipe != nil {
				receiptId = recipe.Id
			}

			sizeCreate := sizefoodmodel.SizeFoodCreate{
				FoodId:   id,
				SizeId:   sizeId,
				Name:     name,
				Cost:     cost,
				Price:    price,
				RecipeId: receiptId,
				Recipe:   recipe,
			}

			////validate data
			if err := sizeCreate.Validate(); err != nil {
				return err
			}

			////define job
			/////for size Food
			sizeFoodJob := asyncjob.NewJob(func(
				sizeFoodCreate sizefoodmodel.SizeFoodCreate,
			) func(ctx context.Context) error {
				return func(ctx context.Context) error {
					return biz.sizeFoodStore.CreateSizeFood(
						ctx,
						&sizeFoodCreate)
				}
			}(sizeCreate))
			sizeFoodJobs = append(sizeFoodJobs, sizeFoodJob)

			/////for recipe
			recipeJob := asyncjob.NewJob(func(
				recipe *recipemodel.RecipeCreate,
			) func(ctx context.Context) error {
				return func(ctx context.Context) error {
					return biz.recipeStore.CreateRecipe(
						ctx,
						recipe)
				}
			}(recipe))
			recipeJobs = append(recipeJobs, recipeJob)

			/////for recipe detail
			recipeDetailJob := asyncjob.NewJob(func(
				recipeDetails []recipedetailmodel.RecipeDetailCreate,
			) func(ctx context.Context) error {
				return func(ctx context.Context) error {
					return biz.recipeDetailStore.CreateListRecipeDetail(
						ctx,
						recipeDetails)
				}
			}(*recipe.Details))
			recipeDetailJobs = append(recipeDetailJobs, recipeDetailJob)
		}

		for key, value := range mapExistSize {
			currentSize := mapSizeFood[key]
			if value == 0 {
				size := mapSizeFoodUpdate[key]
				////define job
				/////for size Food
				sizeFoodJob := asyncjob.NewJob(func(
					foodId string,
					sizeId string,
					sizeFoodUpdate *sizefoodmodel.SizeFoodUpdate,
				) func(ctx context.Context) error {
					return func(ctx context.Context) error {
						return biz.sizeFoodStore.UpdateSizeFood(
							ctx,
							foodId,
							sizeId,
							sizeFoodUpdate)
					}
				}(id, *size.SizeId, &size))
				sizeFoodJobs = append(sizeFoodJobs, sizeFoodJob)

				/////for recipe detail
				if size.Recipe != nil &&
					size.Recipe.Details != nil &&
					len(*size.Recipe.Details) != 0 {
					//////get current receipt detail
					currentReceiptDetails, err := biz.recipeDetailStore.FindListRecipeDetail(
						ctx,
						map[string]interface{}{"recipeId": currentSize.RecipeId})
					if err != nil {
						return err
					}

					//////classify recipe detail
					mapExistRecipeDetail := make(map[string]int)
					mapRecipeDetailUpdate := make(map[string]recipedetailmodel.RecipeDetailUpdate)
					mapRecipeDetail := make(map[string]recipedetailmodel.RecipeDetail)
					for _, v := range *size.Recipe.Details {
						mapExistRecipeDetail[*v.IngredientId]++
						mapRecipeDetailUpdate[*v.IngredientId] = v
					}
					for _, v := range *currentReceiptDetails {
						mapExistRecipeDetail[v.IngredientId]--
						mapRecipeDetail[v.IngredientId] = v
					}

					//////define jobs
					var notExistRecipeDetails []recipedetailmodel.RecipeDetailCreate
					for receiptDetailKey, receiptDetailExistTimes := range mapExistRecipeDetail {
						if receiptDetailExistTimes == -1 {
							currentReceiptDetail := mapRecipeDetail[receiptDetailKey]
							//////job delete recipe detail
							recipeDetailJob := asyncjob.NewJob(func(
								recipeId string,
								ingredientId string,
							) func(ctx context.Context) error {
								return func(ctx context.Context) error {
									return biz.recipeDetailStore.DeleteRecipeDetail(
										ctx,
										map[string]interface{}{
											"recipeId":     recipeId,
											"ingredientId": ingredientId,
										})
								}
							}(currentReceiptDetail.RecipeId,
								currentReceiptDetail.IngredientId,
							))
							recipeDetailJobs = append(recipeDetailJobs, recipeDetailJob)
						}
						if receiptDetailExistTimes == 0 {
							recipeDetailNeedUpdate := mapRecipeDetailUpdate[receiptDetailKey]
							//////job update recipe detail
							recipeDetailJob := asyncjob.NewJob(func(
								recipeId string,
								ingredientId string,
								recipeDetailUpdate *recipedetailmodel.RecipeDetailUpdate,
							) func(ctx context.Context) error {
								return func(ctx context.Context) error {
									return biz.recipeDetailStore.UpdateRecipeDetail(
										ctx,
										recipeId,
										ingredientId,
										recipeDetailUpdate)
								}
							}(
								currentSize.RecipeId,
								*recipeDetailNeedUpdate.IngredientId,
								&recipeDetailNeedUpdate,
							))
							recipeDetailJobs = append(recipeDetailJobs, recipeDetailJob)
						}
						if receiptDetailExistTimes == 1 {
							recipeDetailNeedCreate := mapRecipeDetailUpdate[receiptDetailKey]
							recipeDetailCreate := recipedetailmodel.RecipeDetailCreate{
								RecipeId:     currentSize.RecipeId,
								IngredientId: *recipeDetailNeedCreate.IngredientId,
								AmountNeed:   *recipeDetailNeedCreate.AmountNeed,
							}
							notExistRecipeDetails = append(
								notExistRecipeDetails,
								recipeDetailCreate,
							)
						}
					}

					//////job create recipe detail
					recipeDetailJob := asyncjob.NewJob(func(
						recipeDetails []recipedetailmodel.RecipeDetailCreate,
					) func(ctx context.Context) error {
						return func(ctx context.Context) error {
							return biz.recipeDetailStore.CreateListRecipeDetail(
								ctx,
								recipeDetails)
						}
					}(notExistRecipeDetails))
					recipeDetailJobs = append(recipeDetailJobs, recipeDetailJob)
				}
			}
			if value == -1 {
				////define job
				/////for size food
				sizeFoodJob := asyncjob.NewJob(func(
					foodId string,
					sizeId string,
				) func(ctx context.Context) error {
					return func(ctx context.Context) error {
						return biz.sizeFoodStore.DeleteSizeFood(
							ctx,
							map[string]interface{}{
								"foodId": foodId,
								"sizeId": sizeId,
							})
					}
				}(id, currentSize.SizeId))
				sizeFoodJobs = append(sizeFoodJobs, sizeFoodJob)

				////for recipe
				recipeJob := asyncjob.NewJob(func(
					recipeId string,
				) func(ctx context.Context) error {
					return func(ctx context.Context) error {
						return biz.recipeStore.DeleteRecipe(
							ctx,
							map[string]interface{}{"recipeId": recipeId},
						)
					}
				}(currentSize.RecipeId))
				recipeJobs = append(recipeJobs, recipeJob)

				////for recipe detail
				recipeDetailJob := asyncjob.NewJob(func(
					recipeId string,
				) func(ctx context.Context) error {
					return func(ctx context.Context) error {
						return biz.recipeDetailStore.DeleteRecipeDetail(
							ctx,
							map[string]interface{}{"recipeId": recipeId},
						)
					}
				}(currentSize.RecipeId))
				recipeDetailJobs = append(recipeDetailJobs, recipeDetailJob)
			}
		}
	}

	//handle define job update food
	jobUpdateFood := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.foodStore.UpdateFood(ctx, id, data)
	})

	//combine all job
	jobs := []asyncjob.Job{
		jobUpdateFood,
	}
	jobs = append(jobs, categoryProductJobs...)
	jobs = append(jobs, categoryJobs...)
	jobs = append(jobs, sizeFoodJobs...)
	jobs = append(jobs, recipeJobs...)
	jobs = append(jobs, recipeDetailJobs...)

	//run jobs
	group := asyncjob.NewGroup(
		false,
		jobs...)
	if err := group.Run(context.Background()); err != nil {
		return err
	}

	return nil
}

package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/asyncjob"
	"coffee_shop_management_backend/module/category/categorymodel"
	"coffee_shop_management_backend/module/categoryfood/categoryfoodmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
	"context"
)

type CreateFoodStore interface {
	CreateFood(
		ctx context.Context,
		data *productmodel.FoodCreate,
	) error
}

type CreateCategoryFoodStore interface {
	CreateCategoryFood(
		ctx context.Context,
		data *categoryfoodmodel.CategoryFoodCreate,
	) error
}

type UpdateCategoryStore interface {
	FindCategory(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
	UpdateAmountProductCategory(
		ctx context.Context,
		id string,
		data *categorymodel.CategoryUpdateAmountProduct,
	) error
}

type CreateSizeFoodStore interface {
	CreateSizeFood(
		ctx context.Context,
		data *sizefoodmodel.SizeFoodCreate,
	) error
}

type CreateRecipeStore interface {
	CreateRecipe(
		ctx context.Context,
		data *recipemodel.RecipeCreate,
	) error
}

type CheckIngredientStore interface {
	FindIngredient(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*ingredientmodel.Ingredient, error)
}

type CreateListRecipeDetailStore interface {
	CreateListRecipeDetail(
		ctx context.Context,
		data []recipedetailmodel.RecipeDetailCreate,
	) error
}

type createFoodBiz struct {
	foodStore         CreateFoodStore
	categoryFoodStore CreateCategoryFoodStore
	categoryStore     UpdateCategoryStore
	sizeFoodStore     CreateSizeFoodStore
	recipeStore       CreateRecipeStore
	ingredientStore   CheckIngredientStore
	recipeDetailStore CreateListRecipeDetailStore
}

func NewCreateFoodBiz(
	foodStore CreateFoodStore,
	categoryFoodStore CreateCategoryFoodStore,
	categoryStore UpdateCategoryStore,
	sizeFoodStore CreateSizeFoodStore,
	recipeStore CreateRecipeStore,
	ingredientStore CheckIngredientStore,
	recipeDetailStore CreateListRecipeDetailStore) *createFoodBiz {
	return &createFoodBiz{
		foodStore:         foodStore,
		categoryFoodStore: categoryFoodStore,
		categoryStore:     categoryStore,
		sizeFoodStore:     sizeFoodStore,
		recipeStore:       recipeStore,
		ingredientStore:   ingredientStore,
		recipeDetailStore: recipeDetailStore,
	}
}

func (biz *createFoodBiz) CreateFood(
	ctx context.Context,
	data *productmodel.FoodCreate) error {
	//validate data
	if err := data.Validate(); err != nil {
		return err
	}

	//check category exists
	for _, v := range data.Categories {
		if _, err := biz.categoryStore.FindCategory(
			ctx,
			map[string]interface{}{"id": v}); err != nil {
			return err
		}
	}

	//check ingredients exists
	for _, size := range data.Sizes {
		for _, recipeDetail := range *size.Recipe.Details {
			if _, err := biz.ingredientStore.FindIngredient(
				ctx,
				map[string]interface{}{"id": recipeDetail.IngredientId},
			); err != nil {
				return err
			}
		}
	}

	//handle id food
	idAddress, err := common.IdProcess(data.Id)
	if err != nil {
		return err
	}
	data.Id = idAddress

	//handle id for size food
	for i, _ := range data.Sizes {
		data.Sizes[i].FoodId = *idAddress

		sizeId, err := common.GenerateId()
		if err != nil {
			return err
		}
		data.Sizes[i].SizeId = sizeId

		recipeId, err := common.GenerateId()
		if err != nil {
			return err
		}
		data.Sizes[i].RecipeId = recipeId
		data.Sizes[i].Recipe.Id = recipeId
		for _, recipeDetail := range *data.Sizes[i].Recipe.Details {
			recipeDetailPointer := &recipeDetail
			recipeDetailPointer.RecipeId = recipeId
		}
	}

	//handle store data
	///handle define job create food
	jobCreateFood := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.foodStore.CreateFood(ctx, data)
	})

	///handle define job create category product
	var categoryProductJobs []asyncjob.Job
	for _, value := range data.Categories {
		tempJob := asyncjob.NewJob(func(
			productId string,
			categoryId string,
		) func(ctx context.Context) error {
			return func(ctx context.Context) error {
				categoryFoodCreate := categoryfoodmodel.CategoryFoodCreate{
					FoodId:     &productId,
					CategoryId: &categoryId,
				}
				return biz.categoryFoodStore.CreateCategoryFood(
					ctx,
					&categoryFoodCreate)
			}
		}(*idAddress, value))
		categoryProductJobs = append(categoryProductJobs, tempJob)
	}

	///handle define job update amount product of category
	var categoryJobs []asyncjob.Job
	for _, value := range data.Categories {
		tempJob := asyncjob.NewJob(func(
			productId string,
			categoryId string,
		) func(ctx context.Context) error {
			return func(ctx context.Context) error {
				amount := 1
				categoryUpdateAmountProduct := categorymodel.CategoryUpdateAmountProduct{
					AmountProduct: &amount,
				}
				return biz.categoryStore.UpdateAmountProductCategory(
					ctx,
					categoryId,
					&categoryUpdateAmountProduct)
			}
		}(*idAddress, value))
		categoryJobs = append(categoryJobs, tempJob)
	}

	///handle define job create size product, recipe and recipe details
	var sizeFoodJobs []asyncjob.Job
	var recipeJobs []asyncjob.Job
	var recipeDetailJobs []asyncjob.Job
	for _, value := range data.Sizes {
		////for size Food
		sizeFoodJob := asyncjob.NewJob(func(
			sizeFoodCreate sizefoodmodel.SizeFoodCreate,
		) func(ctx context.Context) error {
			return func(ctx context.Context) error {
				return biz.sizeFoodStore.CreateSizeFood(
					ctx,
					&sizeFoodCreate)
			}
		}(value))
		sizeFoodJobs = append(sizeFoodJobs, sizeFoodJob)

		////for recipe
		recipeOfSize := value.Recipe
		recipeJob := asyncjob.NewJob(func(
			recipe *recipemodel.RecipeCreate,
		) func(ctx context.Context) error {
			return func(ctx context.Context) error {
				return biz.recipeStore.CreateRecipe(
					ctx,
					recipe)
			}
		}(recipeOfSize))
		recipeJobs = append(recipeJobs, recipeJob)

		////for recipe detail
		recipeDetailJob := asyncjob.NewJob(func(
			recipeDetails []recipedetailmodel.RecipeDetailCreate,
		) func(ctx context.Context) error {
			return func(ctx context.Context) error {
				return biz.recipeDetailStore.CreateListRecipeDetail(
					ctx,
					recipeDetails)
			}
		}(*recipeOfSize.Details))
		recipeDetailJobs = append(recipeDetailJobs, recipeDetailJob)
	}

	///combine all job
	jobs := []asyncjob.Job{
		jobCreateFood,
	}
	jobs = append(jobs, categoryProductJobs...)
	jobs = append(jobs, categoryJobs...)
	jobs = append(jobs, recipeJobs...)
	jobs = append(jobs, recipeDetailJobs...)

	///run jobs
	group := asyncjob.NewGroup(
		false,
		jobs...)
	if err := group.Run(context.Background()); err != nil {
		return err
	}
	return nil
}

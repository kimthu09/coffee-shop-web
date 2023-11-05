package ginproduct

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/category/categorystore"
	"coffee_shop_management_backend/module/categoryfood/categoryfoodstore"
	"coffee_shop_management_backend/module/ingredient/ingredientstore"
	"coffee_shop_management_backend/module/product/productbiz"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/product/productrepo"
	"coffee_shop_management_backend/module/product/productstore"
	"coffee_shop_management_backend/module/recipe/recipestore"
	"coffee_shop_management_backend/module/recipedetail/recipedetailstore"
	"coffee_shop_management_backend/module/sizefood/sizefoodstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data productmodel.FoodUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data.IsActive = nil

		db := appCtx.GetMainDBConnection().Begin()

		foodStore := productstore.NewSQLStore(db)
		categoryFoodStore := categoryfoodstore.NewSQLStore(db)
		categoryStore := categorystore.NewSQLStore(db)
		sizeFoodStore := sizefoodstore.NewSQLStore(db)
		recipeStore := recipestore.NewSQLStore(db)
		ingredientStore := ingredientstore.NewSQLStore(db)
		recipeDetailStore := recipedetailstore.NewSQLStore(db)

		repo := productrepo.NewUpdateFoodRepo(
			foodStore,
			categoryFoodStore,
			categoryStore,
			sizeFoodStore,
			recipeStore,
			ingredientStore,
			recipeDetailStore,
		)

		business := productbiz.NewUpdateFoodBiz(repo)

		if err := business.UpdateFood(c.Request.Context(), id, &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}

package ginproduct

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/category/categorystore"
	"coffee_shop_management_backend/module/categoryfood/categoryfoodstore"
	"coffee_shop_management_backend/module/ingredient/ingredientstore"
	"coffee_shop_management_backend/module/product/productbiz"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/product/productstore"
	"coffee_shop_management_backend/module/recipe/recipestore"
	"coffee_shop_management_backend/module/recipedetail/recipedetailstore"
	"coffee_shop_management_backend/module/sizefood/sizefoodstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data productmodel.FoodCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		foodStore := productstore.NewSQLStore(appCtx.GetMainDBConnection())
		categoryFoodStore := categoryfoodstore.NewSQLStore(appCtx.GetMainDBConnection())
		categoryStore := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		sizeFoodStore := sizefoodstore.NewSQLStore(appCtx.GetMainDBConnection())
		recipeStore := recipestore.NewSQLStore(appCtx.GetMainDBConnection())
		ingredientStore := ingredientstore.NewSQLStore(appCtx.GetMainDBConnection())
		recipeDetailStore := recipedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		business := productbiz.NewCreateFoodBiz(
			foodStore,
			categoryFoodStore,
			categoryStore,
			sizeFoodStore,
			recipeStore,
			ingredientStore,
			recipeDetailStore,
		)

		if err := business.CreateFood(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.Id))
	}
}

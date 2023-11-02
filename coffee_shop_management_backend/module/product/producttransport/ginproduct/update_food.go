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

func UpdateFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data productmodel.FoodUpdate

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
		business := productbiz.NewUpdateFoodBiz(
			foodStore,
			categoryFoodStore,
			categoryStore,
			sizeFoodStore,
			recipeStore,
			ingredientStore,
			recipeDetailStore,
		)

		if err := business.UpdateFood(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}

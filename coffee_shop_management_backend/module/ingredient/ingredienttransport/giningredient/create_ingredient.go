package giningredient

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/ingredient/ingredientbiz"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateIngredient(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data ingredientmodel.IngredientCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		gen := generator.NewShortIdGenerator()
		store := ingredientstore.NewSQLStore(appCtx.GetMainDBConnection())
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		business := ingredientbiz.NewCreateIngredientBiz(gen, store, requester)

		if err := business.CreateIngredient(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.Id))
	}
}

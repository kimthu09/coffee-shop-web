package giningredientdetail

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailbiz"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListIngredientDetailById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ingredientId := c.Param("id")

		var filter ingredientdetailmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := ingredientdetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := ingredientdetailbiz.NewListIngredientDetailByIdBiz(store)

		result, err := biz.ListIngredientDetailById(c.Request.Context(), ingredientId, &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

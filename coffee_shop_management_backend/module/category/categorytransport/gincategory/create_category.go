package gincategory

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/category/categorybiz"
	"coffee_shop_management_backend/module/category/categorymodel"
	"coffee_shop_management_backend/module/category/categorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data categorymodel.CategoryCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		gen := generator.NewShortIdGenerator()
		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		business := categorybiz.NewCreateCategoryBiz(gen, store, requester)

		if err := business.CreateCategory(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}

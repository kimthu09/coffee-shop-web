package gincategory

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/category/categorybiz"
	"coffee_shop_management_backend/module/category/categorymodel"
	"coffee_shop_management_backend/module/category/categorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter categorymodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := categorybiz.NewListCategoryBiz(store, requester)

		result, err := biz.ListCategory(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

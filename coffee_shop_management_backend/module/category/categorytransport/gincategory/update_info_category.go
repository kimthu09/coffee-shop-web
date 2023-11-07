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

func UpdateInfoCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data categorymodel.CategoryUpdateInfo

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := categorybiz.NewUpdateInfoCategoryBiz(store, requester)

		if err := biz.UpdateInfoCategory(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}

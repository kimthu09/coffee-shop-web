package ginproduct

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productbiz"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/product/productrepo"
	"coffee_shop_management_backend/module/product/productstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeStatusToppings(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data []productmodel.ToppingUpdateStatus

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := productstore.NewSQLStore(db)
		repo := productrepo.NewChangeStatusToppingsRepo(store)
		business := productbiz.NewChangeStatusToppingsBiz(repo, requester)

		if err := business.ChangeStatusToppings(c.Request.Context(), data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

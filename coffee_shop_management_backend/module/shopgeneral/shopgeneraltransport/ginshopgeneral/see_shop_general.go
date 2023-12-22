package ginshopgeneral

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/shopgeneral/shopgeneralbiz"
	"coffee_shop_management_backend/module/shopgeneral/shopgeneralstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeShopGeneral(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection()

		store := shopgeneralstore.NewSQLStore(db)

		business := shopgeneralbiz.NewSeeShopGeneralBiz(
			store,
			requester,
		)

		general, err := business.SeeShopGeneral(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(general))
	}
}

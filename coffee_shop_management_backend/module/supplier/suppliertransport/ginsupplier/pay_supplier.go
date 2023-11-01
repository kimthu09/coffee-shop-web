package ginsupplier

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/supplier/supplierbiz"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplier/supplierstore"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PaySupplier(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data suppliermodel.SupplierUpdateDebt

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(common.Requester)

		supplierStore := supplierstore.NewSQLStore(appCtx.GetMainDBConnection())
		supplierDebtStore := supplierdebtstore.NewSQLStore(appCtx.GetMainDBConnection())
		business := supplierbiz.NewUpdatePayBiz(supplierStore, supplierDebtStore, requester)

		idSupplierDebt, err := business.PaySupplier(c.Request.Context(), id, &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(idSupplierDebt))
	}
}

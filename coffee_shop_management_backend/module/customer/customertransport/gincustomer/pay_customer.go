package gincustomer

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/customer/customerbiz"
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/customer/customerstore"
	"coffee_shop_management_backend/module/customerdebt/customerdebtstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PayCustomer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data customermodel.CustomerUpdateDebt

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(common.Requester)

		customerStore := customerstore.NewSQLStore(appCtx.GetMainDBConnection())
		customerDebtStore := customerdebtstore.NewSQLStore(appCtx.GetMainDBConnection())
		business := customerbiz.NewUpdatePayBiz(customerStore, customerDebtStore, requester)

		idCustomerDebt, err := business.PayCustomer(c.Request.Context(), id, &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(idCustomerDebt))
	}
}

package gincustomer

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/customer/customerbiz"
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/customer/customerstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCustomer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data customermodel.CustomerCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := customerstore.NewSQLStore(appCtx.GetMainDBConnection())
		business := customerbiz.NewCreateCustomerBiz(store)

		if err := business.CreateCustomer(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.Id))
	}
}

package gincustomer

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customerbiz"
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/customer/customerrepo"
	"coffee_shop_management_backend/module/customer/customerstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCustomer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter customermodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := customerstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := customerrepo.NewListCustomerRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := customerbiz.NewListCustomerBiz(repo, requester)

		result, err := biz.ListCustomer(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

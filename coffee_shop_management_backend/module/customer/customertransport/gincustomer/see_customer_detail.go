package gincustomer

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customerbiz"
	"coffee_shop_management_backend/module/customer/customerrepo"
	"coffee_shop_management_backend/module/customer/customerstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeCustomerDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		customerStore := customerstore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := customerrepo.NewSeeCustomerDetailRepo(customerStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := customerbiz.NewSeeCustomerDetailBiz(repo, requester)

		result, err := biz.SeeCustomerDetail(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}

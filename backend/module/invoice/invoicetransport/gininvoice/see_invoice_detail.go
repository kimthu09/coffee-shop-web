package gininvoice

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/invoice/invoicebiz"
	"coffee_shop_management_backend/module/invoice/invoicerepo"
	"coffee_shop_management_backend/module/invoice/invoicestore"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeInvoiceDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		invoiceDetailStore := invoicedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		invoiceStore := invoicestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := invoicerepo.NewSeeInvoiceDetailRepo(invoiceDetailStore, invoiceStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := invoicebiz.NewSeeInvoiceDetailBiz(
			repo, requester)

		result, err := biz.SeeInvoiceDetail(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}

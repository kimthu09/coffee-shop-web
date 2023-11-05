package gininvoice

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/customer/customerstore"
	"coffee_shop_management_backend/module/customerdebt/customerdebtstore"
	"coffee_shop_management_backend/module/invoice/invoicebiz"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"coffee_shop_management_backend/module/invoice/invoicerepo"
	"coffee_shop_management_backend/module/invoice/invoicestore"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailstore"
	"coffee_shop_management_backend/module/product/productstore"
	"coffee_shop_management_backend/module/sizefood/sizefoodstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateInvoice(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data invoicemodel.InvoiceCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(common.Requester)
		data.CreateBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		invoiceStore := invoicestore.NewSQLStore(db)
		invoiceDetailStore := invoicedetailstore.NewSQLStore(db)
		customerStore := customerstore.NewSQLStore(db)
		customerDebtStore := customerdebtstore.NewSQLStore(db)
		sizeFoodStore := sizefoodstore.NewSQLStore(db)
		foodStore := productstore.NewSQLStore(db)
		toppingStore := productstore.NewSQLStore(db)

		repo := invoicerepo.NewCreateInvoiceRepo(
			invoiceStore,
			invoiceDetailStore,
			customerStore,
			customerDebtStore,
			sizeFoodStore,
			foodStore,
			toppingStore,
		)

		business := invoicebiz.NewCreateInvoiceBiz(repo)

		if err := business.CreateInvoice(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.Id))
	}
}

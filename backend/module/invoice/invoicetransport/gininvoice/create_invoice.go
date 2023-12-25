package gininvoice

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customerstore"
	"coffee_shop_management_backend/module/ingredient/ingredientstore"
	"coffee_shop_management_backend/module/invoice/invoicebiz"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"coffee_shop_management_backend/module/invoice/invoicerepo"
	"coffee_shop_management_backend/module/invoice/invoicestore"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailstore"
	"coffee_shop_management_backend/module/product/productstore"
	"coffee_shop_management_backend/module/shopgeneral/shopgeneralstore"
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

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreatedBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		invoiceStore := invoicestore.NewSQLStore(db)
		invoiceDetailStore := invoicedetailstore.NewSQLStore(db)
		customerStore := customerstore.NewSQLStore(db)
		sizeFoodStore := sizefoodstore.NewSQLStore(db)
		foodStore := productstore.NewSQLStore(db)
		toppingStore := productstore.NewSQLStore(db)
		ingredientStore := ingredientstore.NewSQLStore(db)
		shopGeneralStore := shopgeneralstore.NewSQLStore(db)

		repo := invoicerepo.NewCreateInvoiceRepo(
			invoiceStore,
			invoiceDetailStore,
			customerStore,
			sizeFoodStore,
			foodStore,
			toppingStore,
			ingredientStore,
			shopGeneralStore,
		)

		gen := generator.NewShortIdGenerator()

		business := invoicebiz.NewCreateInvoiceBiz(gen, repo, requester)

		if err := business.CreateInvoice(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"id":           data.Id,
			"customer":     data.Customer,
			"shopName":     data.ShopName,
			"shopPhone":    data.ShopPhone,
			"shopAddress":  data.ShopAddress,
			"shopPassWifi": data.ShopPassWifi,
			"details":      data.InvoiceDetails,
			"total":        data.TotalPrice,
			"received":     data.AmountReceived,
			"discount":     data.AmountPriceUsePoint,
		})
	}
}

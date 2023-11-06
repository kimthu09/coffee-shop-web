package ginsupplier

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/supplier/supplierbiz"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplier/supplierrepo"
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

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreateBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		supplierStore := supplierstore.NewSQLStore(db)
		supplierDebtStore := supplierdebtstore.NewSQLStore(db)
		repo := supplierrepo.NewUpdatePayRepo(supplierStore, supplierDebtStore)

		gen := generator.NewShortIdGenerator()

		business := supplierbiz.NewUpdatePayBiz(gen, repo, requester)

		idSupplierDebt, err := business.PaySupplier(c.Request.Context(), id, &data)

		if err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(idSupplierDebt))
	}
}

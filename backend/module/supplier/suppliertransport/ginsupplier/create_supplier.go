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
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSupplier(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data suppliermodel.SupplierCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := supplierstore.NewSQLStore(db)
		repo := supplierrepo.NewCreateSupplierRepo(store)

		gen := generator.NewShortIdGenerator()

		business := supplierbiz.NewCreateSupplierBiz(gen, repo, requester)

		if err := business.CreateSupplier(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}

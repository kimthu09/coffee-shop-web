package ginsupplier

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/supplier/supplierbiz"
	"coffee_shop_management_backend/module/supplier/supplierrepo"
	"coffee_shop_management_backend/module/supplier/supplierstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllSupplier(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := supplierstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := supplierrepo.NewGetAllSupplierRepo(store)

		biz := supplierbiz.GetAllSupplierRepo(repo)

		result, err := biz.GetAllSupplier(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}

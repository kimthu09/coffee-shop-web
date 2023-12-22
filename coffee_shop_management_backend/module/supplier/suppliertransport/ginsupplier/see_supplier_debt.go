package ginsupplier

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/supplier/supplierbiz"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"coffee_shop_management_backend/module/supplier/supplierrepo"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeSupplierDebt(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var debtSupplierFilter filter.SupplierDebtFilter
		if err := c.ShouldBind(&debtSupplierFilter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		supplierDebtStore := supplierdebtstore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := supplierrepo.NewSeeSupplierDebtRepo(supplierDebtStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := supplierbiz.NewSeeSupplierDebtBiz(repo, requester)

		result, err := biz.SeeSupplierDebt(
			c.Request.Context(), id, &debtSupplierFilter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}

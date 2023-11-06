package gincustomer

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customerbiz"
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/customer/customerrepo"
	"coffee_shop_management_backend/module/customer/customerstore"
	"coffee_shop_management_backend/module/customerdebt/customerdebtstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PayCustomer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data customermodel.CustomerUpdateDebt

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreateBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		customerStore := customerstore.NewSQLStore(db)
		customerDebtStore := customerdebtstore.NewSQLStore(db)

		repo := customerrepo.NewUpdatePayRepo(customerStore, customerDebtStore)

		gen := generator.NewShortIdGenerator()

		business := customerbiz.NewUpdatePayBiz(gen, repo, requester)

		customerDebtId, err := business.PayCustomer(c.Request.Context(), id, &data)

		if err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(customerDebtId))
	}
}

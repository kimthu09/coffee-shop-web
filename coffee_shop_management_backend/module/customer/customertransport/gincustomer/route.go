package gincustomer

import (
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	customers := router.Group("/customers", middleware.RequireAuth(appCtx))
	{
		customers.GET("", ListCustomer(appCtx))
		customers.POST("", CreateCustomer(appCtx))
		customers.GET("/:id", SeeCustomerDetail(appCtx))
		customers.GET("/:id/invoices", SeeCustomerInvoice(appCtx))
		customers.POST("/:id", UpdateInfoCustomer(appCtx))
	}
}

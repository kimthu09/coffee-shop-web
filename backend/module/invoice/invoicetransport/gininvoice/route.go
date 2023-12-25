package gininvoice

import (
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	invoices := router.Group("/invoices", middleware.RequireAuth(appCtx))
	{
		invoices.GET("", ListInvoice(appCtx))
		invoices.POST("", CreateInvoice(appCtx))
		invoices.GET("/:id", SeeInvoiceDetail(appCtx))
	}
}

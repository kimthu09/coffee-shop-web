package ginsupplier

import (
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	suppliers := router.Group("/suppliers", middleware.RequireAuth(appCtx))
	{
		suppliers.GET("", ListSupplier(appCtx))
		suppliers.GET("/all", GetAllSupplier(appCtx))
		suppliers.POST("", CreateSupplier(appCtx))
		suppliers.POST("/pay", PaySupplier(appCtx))
		suppliers.GET("/:id", SeeSupplierDetail(appCtx))
		suppliers.GET("/:id/import", SeeSupplierImportNote(appCtx))
		suppliers.GET("/:id/debt", SeeSupplierDebt(appCtx))
		suppliers.POST("/:id", UpdateInfoSupplier(appCtx))
	}
}

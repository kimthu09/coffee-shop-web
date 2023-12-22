package gincategory

import (
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	categories := router.Group("/categories", middleware.RequireAuth(appCtx))
	{
		categories.GET("", ListCategory(appCtx))
		categories.POST("", CreateCategory(appCtx))
		categories.POST("/:id", UpdateInfoCategory(appCtx))
	}
}

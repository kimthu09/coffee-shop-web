package giningredient

import (
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	ingredients := router.Group("/ingredients", middleware.RequireAuth(appCtx))
	{
		ingredients.GET("", ListIngredient(appCtx))
		ingredients.POST("", CreateIngredient(appCtx))
	}
}

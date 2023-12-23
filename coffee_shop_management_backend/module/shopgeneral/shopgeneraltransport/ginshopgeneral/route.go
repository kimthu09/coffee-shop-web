package ginshopgeneral

import (
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	shopGenerals := router.Group("/shop", middleware.RequireAuth(appCtx))
	{
		shopGenerals.GET("", SeeShopGeneral(appCtx))
		shopGenerals.PATCH("", UpdateShopGeneral(appCtx))
	}
}

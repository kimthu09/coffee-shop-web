package ginshopgeneral

import (
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	shopGenerals := router.Group("/shopGenerals", middleware.RequireAuth(appCtx))
	{
		shopGenerals.GET("", SeeShopGeneral(appCtx))
		shopGenerals.POST("", UpdateShopGeneral(appCtx))
	}
}

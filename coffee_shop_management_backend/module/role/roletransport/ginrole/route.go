package ginrole

import (
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	roles := router.Group("/roles", middleware.RequireAuth(appCtx))
	{
		roles.GET("", ListRole(appCtx))
		roles.POST("", CreateRole(appCtx))
		roles.GET("/:id", SeeDetailRole(appCtx))
		roles.POST("/:id", UpdateRole(appCtx))
	}
}

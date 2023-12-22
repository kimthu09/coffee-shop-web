package ginproduct

import (
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	foods := router.Group("/foods", middleware.RequireAuth(appCtx))
	{
		foods.GET("", ListFood(appCtx))
		foods.POST("", CreateFood(appCtx))
		foods.GET("/:id", SeeDetailFood(appCtx))
		foods.POST("/:id", UpdateFood(appCtx))
		foods.POST("/status", ChangeStatusFoods(appCtx))
	}
	toppings := router.Group("/toppings", middleware.RequireAuth(appCtx))
	{
		toppings.GET("", ListTopping(appCtx))
		toppings.POST("", CreateTopping(appCtx))
		toppings.GET("/:id", SeeDetailTopping(appCtx))
		toppings.POST("/:id", UpdateTopping(appCtx))
		toppings.POST("/status", ChangeStatusToppings(appCtx))
	}
}

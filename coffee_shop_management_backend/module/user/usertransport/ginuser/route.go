package ginuser

import (
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	router.POST("/login", Login(appCtx))
	router.POST("/refreshToken", RefreshToken(appCtx))
	auth := router.Group("", middleware.RequireAuth(appCtx))
	{
		auth.GET("/profile", SeeProfile(appCtx), middleware.RequireAuth(appCtx))
	}
	users := router.Group("/users", middleware.RequireAuth(appCtx))
	{
		users.GET("", ListUser(appCtx))
		users.GET("/all", GetAllUser(appCtx))
		users.GET("/:id", SeeUserDetail(appCtx))
		users.POST("", CreateUser(appCtx))
		users.PATCH("/:id/info", UpdateInfoUser(appCtx))
		users.PATCH("/status", ChangeStatusUsers(appCtx))
		users.PATCH("/:id/role", ChangeRoleUser(appCtx))
		users.PATCH("/:id/reset", ResetPassword(appCtx))
		users.PATCH("/:id/password", UpdatePassword(appCtx))
	}
}

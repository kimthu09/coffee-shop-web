package main

import (
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/cancelnote/cancelnotetransport/gincancelnote"
	"coffee_shop_management_backend/module/category/categorytransport/gincategory"
	"coffee_shop_management_backend/module/customer/customertransport/gincustomer"
	"coffee_shop_management_backend/module/exportnote/exportnotetransport/ginexportnote"
	"coffee_shop_management_backend/module/importnote/importnotetransport/ginimportnote"
	"coffee_shop_management_backend/module/ingredient/ingredienttransport/giningredient"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailtransport/giningredientdetail"
	"coffee_shop_management_backend/module/product/producttransport/ginproduct"
	"coffee_shop_management_backend/module/supplier/suppliertransport/ginsupplier"
	"coffee_shop_management_backend/module/user/usertransport/gin_user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	appCtx := appctx.NewAppContext(db, secretKey)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")
	{
		v1.POST("/register", gin_user.Register(appCtx))
		v1.POST("/login", gin_user.Login(appCtx))
	}

	categories := v1.Group("/categories", middleware.RequireAuth(appCtx))
	{
		categories.POST("", gincategory.CreateCategory(appCtx))
		categories.GET("", gincategory.ListCategory(appCtx))
		categories.PATCH("/:id", gincategory.UpdateInfoCategory(appCtx))
	}

	toppings := v1.Group("/toppings", middleware.RequireAuth(appCtx))
	{
		toppings.POST("", ginproduct.CreateTopping(appCtx))
		toppings.PATCH("/:id", ginproduct.UpdateTopping(appCtx))
		toppings.PATCH("/:id/status", ginproduct.ChangeStatusTopping(appCtx))
	}

	foods := v1.Group("/foods", middleware.RequireAuth(appCtx))
	{
		foods.POST("", ginproduct.CreateFood(appCtx))
		foods.PATCH("/:id", ginproduct.UpdateFood(appCtx))
		foods.PATCH("/:id/status", ginproduct.ChangeStatusFood(appCtx))
	}

	suppliers := v1.Group("/suppliers", middleware.RequireAuth(appCtx))
	{
		suppliers.POST("", ginsupplier.CreateSupplier(appCtx))
		suppliers.PATCH("/:id", ginsupplier.UpdateInfoSupplier(appCtx))
		suppliers.POST("/:id/pay", ginsupplier.PaySupplier(appCtx))
	}

	customers := v1.Group("/customers", middleware.RequireAuth(appCtx))
	{
		customers.POST("", gincustomer.CreateCustomer(appCtx))
		customers.PATCH("/:id", gincustomer.UpdateInfoCustomer(appCtx))
		customers.POST("/:id/pay", gincustomer.PayCustomer(appCtx))
	}

	ingredients := v1.Group("/ingredients", middleware.RequireAuth(appCtx))
	{
		ingredients.POST("", giningredient.CreateIngredient(appCtx))
		ingredients.GET("/:id/details", giningredientdetail.ListIngredientDetailById(appCtx))
	}

	importNotes := v1.Group("/importNotes", middleware.RequireAuth(appCtx))
	{
		importNotes.POST("", ginimportnote.CreateImportNote(appCtx))
		importNotes.PATCH("/:id", ginimportnote.ChangeStatusImportNote(appCtx))
	}

	exportNotes := v1.Group("/exportNotes", middleware.RequireAuth(appCtx))
	{
		exportNotes.POST("", ginexportnote.CreateExportNote(appCtx))
	}

	cancelNotes := v1.Group("/cancelNotes", middleware.RequireAuth(appCtx))
	{
		cancelNotes.POST("", gincancelnote.CreateCancelNote(appCtx))
	}

	err = r.Run(":8080")
	if err != nil {
		return
	}
}

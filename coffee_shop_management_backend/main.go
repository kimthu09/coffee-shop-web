package main

import (
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/cancelnote/cancelnotetransport/gincancelnote"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailtransport/gincancelnotedetail"
	"coffee_shop_management_backend/module/category/categorytransport/gincategory"
	"coffee_shop_management_backend/module/customer/customertransport/gincustomer"
	"coffee_shop_management_backend/module/customerdebt/customerdebttransport/gincustomerdebt"
	"coffee_shop_management_backend/module/exportnote/exportnotetransport/ginexportnote"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailtransport/ginexportnotedetail"
	"coffee_shop_management_backend/module/feature/featuretransport/ginfeature"
	"coffee_shop_management_backend/module/importnote/importnotetransport/ginimportnote"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailtransport/ginimportnotedetail"
	"coffee_shop_management_backend/module/ingredient/ingredienttransport/giningredient"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailtransport/giningredientdetail"
	"coffee_shop_management_backend/module/invoice/invoicetransport/gininvoice"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailtransport/gininvoicedetail"
	"coffee_shop_management_backend/module/product/producttransport/ginproduct"
	"coffee_shop_management_backend/module/role/roletransport/ginrole"
	"coffee_shop_management_backend/module/supplier/suppliertransport/ginsupplier"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebttransport/ginsupplierdebt"
	"coffee_shop_management_backend/module/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	//dsn := os.Getenv("DBConnectionStr")
	//secretKey := os.Getenv("SYSTEM_SECRET")

	dsn := "root:123456@tcp(127.0.0.1:33062)/coffeemanagement?charset=utf8mb4&parseTime=True&loc=Local"
	secretKey := "123456789"

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
		v1.POST("/login", ginuser.Login(appCtx))
	}

	users := v1.Group("/users", middleware.RequireAuth(appCtx))
	{
		users.GET("", ginuser.ListUser(appCtx))
		users.PATCH("/:id/info", ginuser.UpdateInfoUser(appCtx))
		users.PATCH("/:id/status", ginuser.ChangeStatusUser(appCtx))
		users.PATCH("/:id/role", ginuser.ChangeRoleUser(appCtx))
		users.PATCH("/:id/reset", ginuser.ResetPassword(appCtx))
	}

	profile := v1.Group("/profile", middleware.RequireAuth(appCtx))
	{
		profile.PATCH("/password", ginuser.UpdatePassword(appCtx))
	}

	categories := v1.Group("/categories", middleware.RequireAuth(appCtx))
	{
		categories.GET("", gincategory.ListCategory(appCtx))
		categories.POST("", gincategory.CreateCategory(appCtx))
		categories.PATCH("/:id", gincategory.UpdateInfoCategory(appCtx))
	}

	toppings := v1.Group("/toppings", middleware.RequireAuth(appCtx))
	{
		toppings.GET("", ginproduct.ListTopping(appCtx))
		toppings.GET("/:id", ginproduct.SeeDetailTopping(appCtx))
		toppings.POST("", ginproduct.CreateTopping(appCtx))
		toppings.PATCH("/:id", ginproduct.UpdateTopping(appCtx))
		toppings.PATCH("/:id/status", ginproduct.ChangeStatusTopping(appCtx))
	}

	foods := v1.Group("/foods", middleware.RequireAuth(appCtx))
	{
		foods.GET("", ginproduct.ListFood(appCtx))
		foods.GET("/:id", ginproduct.SeeDetailFood(appCtx))
		foods.POST("", ginproduct.CreateFood(appCtx))
		foods.PATCH("/:id", ginproduct.UpdateFood(appCtx))
		foods.PATCH("/:id/status", ginproduct.ChangeStatusFood(appCtx))
	}

	suppliers := v1.Group("/suppliers", middleware.RequireAuth(appCtx))
	{
		suppliers.GET("", ginsupplier.ListSupplier(appCtx))
		suppliers.GET("/:id", ginsupplierdebt.ListSupplierDebt(appCtx))
		suppliers.POST("", ginsupplier.CreateSupplier(appCtx))
		suppliers.PATCH("/:id", ginsupplier.UpdateInfoSupplier(appCtx))
		suppliers.POST("/:id/pay", ginsupplier.PaySupplier(appCtx))
	}

	customers := v1.Group("/customers", middleware.RequireAuth(appCtx))
	{
		customers.GET("", gincustomer.ListCustomer(appCtx))
		customers.GET("/:id", gincustomerdebt.ListCustomerDebt(appCtx))
		customers.POST("", gincustomer.CreateCustomer(appCtx))
		customers.PATCH("/:id", gincustomer.UpdateInfoCustomer(appCtx))
		customers.POST("/:id/pay", gincustomer.PayCustomer(appCtx))
	}

	ingredients := v1.Group("/ingredients", middleware.RequireAuth(appCtx))
	{
		ingredients.GET("", giningredient.ListIngredient(appCtx))
		ingredients.GET("/:id", giningredientdetail.ListIngredientDetail(appCtx))
		ingredients.POST("", giningredient.CreateIngredient(appCtx))
	}

	importNotes := v1.Group("/importNotes", middleware.RequireAuth(appCtx))
	{
		importNotes.GET("", ginimportnote.ListImportNote(appCtx))
		importNotes.GET("/:id", ginimportnotedetail.ListImportNoteDetail(appCtx))
		importNotes.POST("", ginimportnote.CreateImportNote(appCtx))
		importNotes.PATCH("/:id", ginimportnote.ChangeStatusImportNote(appCtx))
	}

	exportNotes := v1.Group("/exportNotes", middleware.RequireAuth(appCtx))
	{
		exportNotes.GET("", ginexportnote.ListExportNote(appCtx))
		exportNotes.GET("/:id", ginexportnotedetail.ListExportNoteDetail(appCtx))
		exportNotes.POST("", ginexportnote.CreateExportNote(appCtx))
	}

	cancelNotes := v1.Group("/cancelNotes", middleware.RequireAuth(appCtx))
	{
		cancelNotes.GET("", gincancelnote.ListCancelNote(appCtx))
		cancelNotes.GET("/:id", gincancelnotedetail.ListCancelNoteDetail(appCtx))
		cancelNotes.POST("", gincancelnote.CreateCancelNote(appCtx))
	}

	invoices := v1.Group("/invoices", middleware.RequireAuth(appCtx))
	{
		invoices.GET("", gininvoice.ListInvoice(appCtx))
		invoices.GET("/:id", gininvoicedetail.ListInvoiceDetail(appCtx))
		invoices.POST("", gininvoice.CreateInvoice(appCtx))
	}

	roles := v1.Group("/roles", middleware.RequireAuth(appCtx))
	{
		roles.GET("", ginrole.ListRole(appCtx))
		roles.POST("", ginrole.CreateRole(appCtx))
		roles.PATCH("/:id", ginrole.UpdateRole(appCtx))
	}

	features := v1.Group("/features", middleware.RequireAuth(appCtx))
	{
		features.GET("", ginfeature.ListFeature(appCtx))
	}

	err = r.Run(":8080")
	if err != nil {
		return
	}
}

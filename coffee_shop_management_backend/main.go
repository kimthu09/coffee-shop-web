package main

import (
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/category/categorytransport/gincategory"
	"coffee_shop_management_backend/module/customer/customertransport/gincustomer"
	"coffee_shop_management_backend/module/exportnote/exportnotetransport/ginexportnote"
	"coffee_shop_management_backend/module/feature/featuretransport/ginfeature"
	"coffee_shop_management_backend/module/importnote/importnotetransport/ginimportnote"
	"coffee_shop_management_backend/module/ingredient/ingredienttransport/giningredient"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotetransport/gininventorychecknote"
	"coffee_shop_management_backend/module/invoice/invoicetransport/gininvoice"
	"coffee_shop_management_backend/module/product/producttransport/ginproduct"
	"coffee_shop_management_backend/module/role/roletransport/ginrole"
	"coffee_shop_management_backend/module/shopgeneral/shopgeneraltransport/ginshopgeneral"
	"coffee_shop_management_backend/module/supplier/suppliertransport/ginsupplier"
	"coffee_shop_management_backend/module/user/usertransport/ginuser"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type appConfig struct {
	Port string
	Env  string

	DBConnStr string

	SecretKey string
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln("Error when loading config:", err)
	}

	fmt.Println("Connecting to database...")
	db, err := connectDatabaseWithRetryIn60s(cfg)
	if err != nil {
		log.Fatalln("Error when connecting to database:", err)
	}

	if cfg.Env == "dev" {
		db = db.Debug()
	}

	appCtx := appctx.NewAppContext(db, cfg.SecretKey)

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")
	{
		gincategory.SetupRoutes(v1, appCtx)
		gincustomer.SetupRoutes(v1, appCtx)
		ginexportnote.SetupRoutes(v1, appCtx)
		ginfeature.SetupRoutes(v1, appCtx)
		gininvoice.SetupRoutes(v1, appCtx)
		ginimportnote.SetupRoutes(v1, appCtx)
		giningredient.SetupRoutes(v1, appCtx)
		gininventorychecknote.SetupRoutes(v1, appCtx)
		ginproduct.SetupRoutes(v1, appCtx)
		ginrole.SetupRoutes(v1, appCtx)
		ginshopgeneral.SetupRoutes(v1, appCtx)
		ginsupplier.SetupRoutes(v1, appCtx)
		ginuser.SetupRoutes(v1, appCtx)
	}

	if err := r.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatalln("Error running server:", err)
	}
}

func loadConfig() (*appConfig, error) {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatalln("Error when loading .env", err)
	}

	return &appConfig{
		Port:      env["PORT"],
		Env:       env["ENVIRONMENT"],
		DBConnStr: env["DB_CONNECTION_STR"],
		SecretKey: env["SECRET_KEY"],
	}, nil
}

func connectDatabaseWithRetryIn60s(cfg *appConfig) (*gorm.DB, error) {
	const timeRetry = 60 * time.Second
	var connectDatabase = func(cfg *appConfig) (*gorm.DB, error) {
		db, err := gorm.Open(mysql.Open(cfg.DBConnStr), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return db.Debug(), nil
	}

	var db *gorm.DB
	var err error

	deadline := time.Now().Add(timeRetry)

	for time.Now().Before(deadline) {
		log.Println("Connecting to database...")
		db, err = connectDatabase(cfg)
		if err == nil {
			return db, nil
		}
		time.Sleep(time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after retrying for 10 seconds: %w", err)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

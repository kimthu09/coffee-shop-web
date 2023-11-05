package ginproduct

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/product/productbiz"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/product/productrepo"
	"coffee_shop_management_backend/module/product/productstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeStatusTopping(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data productmodel.ToppingUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if data.IsActive == nil {
			panic(common.ErrInvalidRequest(productmodel.ErrProductIsActiveEmpty))
		}
		data.Name = nil
		data.Description = nil
		data.CookingGuide = nil
		data.Recipe = nil

		db := appCtx.GetMainDBConnection().Begin()

		store := productstore.NewSQLStore(db)
		repo := productrepo.NewChangeStatusToppingRepo(store)
		business := productbiz.NewChangeStatusToppingBiz(repo)

		if err := business.ChangeStatusTopping(c.Request.Context(), id, &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}

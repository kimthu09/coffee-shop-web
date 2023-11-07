package ginproduct

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productbiz"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/product/productrepo"
	"coffee_shop_management_backend/module/product/productstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeStatusFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data productmodel.FoodUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if data.IsActive == nil {
			panic(common.ErrInvalidRequest(productmodel.ErrProductIsActiveEmpty))
		}

		data.Name = nil
		data.Description = nil
		data.CookingGuide = nil
		data.Categories = nil
		data.Sizes = nil

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := productstore.NewSQLStore(db)
		repo := productrepo.NewChangeStatusFoodRepo(store)
		business := productbiz.NewChangeStatusFoodBiz(repo, requester)

		if err := business.ChangeStatusFood(c.Request.Context(), id, &data); err != nil {
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

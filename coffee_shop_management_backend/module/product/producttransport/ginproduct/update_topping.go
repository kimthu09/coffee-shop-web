package ginproduct

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productbiz"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/product/productrepo"
	"coffee_shop_management_backend/module/product/productstore"
	"coffee_shop_management_backend/module/recipedetail/recipedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateTopping(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data productmodel.ToppingUpdateInfo

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		toppingStore := productstore.NewSQLStore(db)
		recipeDetailStore := recipedetailstore.NewSQLStore(db)

		repo := productrepo.NewUpdateToppingRepo(
			toppingStore,
			recipeDetailStore,
		)

		biz := productbiz.NewUpdateToppingBiz(repo, requester)

		if err := biz.UpdateTopping(c.Request.Context(), id, &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

package ginproduct

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productbiz"
	"coffee_shop_management_backend/module/product/productrepo"
	"coffee_shop_management_backend/module/product/productstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeDetailFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		store := productstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := productrepo.NewSeeDetailFoodRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := productbiz.NewSeeDetailFoodBiz(repo, requester)

		result, err := biz.SeeDetailFood(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))
	}
}

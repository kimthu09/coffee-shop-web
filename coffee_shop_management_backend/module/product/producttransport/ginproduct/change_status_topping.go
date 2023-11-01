package ginproduct

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/product/productbiz"
	"coffee_shop_management_backend/module/product/productstore"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ChangeStatusTopping(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		isActive, err := strconv.ParseBool(c.Param("isActive"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := productstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewChangeStatusToppingBiz(store)

		if err := biz.ChangeStatusTopping(c.Request.Context(), id, isActive); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}

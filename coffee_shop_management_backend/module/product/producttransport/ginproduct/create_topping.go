package ginproduct

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/product/productbiz"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/product/productstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTopping(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data productmodel.ToppingCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := productstore.NewSQLStore(appCtx.GetMainDBConnection())
		business := productbiz.NewCreateToppingBiz(store)

		if err := business.CreateTopping(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.Id))
	}
}

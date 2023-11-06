package gincancelnotedetail

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailbiz"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListIngredientDetailById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		cancelNoteId := c.Param("id")

		var filter cancelnotedetailmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := cancelnotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := cancelnotedetailbiz.NewListCancelNoteDetailByIdBiz(store)

		result, err := biz.ListCancelNoteDetailByIdBiz(c.Request.Context(), cancelNoteId, &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

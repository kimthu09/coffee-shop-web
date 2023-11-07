package gincancelnotedetail

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailbiz"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCancelNoteDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := cancelnotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := cancelnotedetailbiz.NewListCancelNoteDetailBiz(store, requester)

		result, err := biz.ListCancelNoteDetail(c.Request.Context(), id, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}

package ginimportnotedetail

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailbiz"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListImportNoteDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := importnotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := importnotedetailbiz.NewListImportNoteDetailBiz(store, requester)

		result, err := biz.ListImportNoteDetail(c.Request.Context(), id, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}

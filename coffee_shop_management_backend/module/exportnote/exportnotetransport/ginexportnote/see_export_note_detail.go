package ginexportnote

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/exportnote/exportnotebiz"
	"coffee_shop_management_backend/module/exportnote/exportnoterepo"
	"coffee_shop_management_backend/module/exportnote/exportnotestore"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeExportNoteDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		exportNoteDetailStore := exportnotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		exportNoteStore := exportnotestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := exportnoterepo.NewSeeExportNoteDetailRepo(exportNoteDetailStore, exportNoteStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := exportnotebiz.NewSeeExportNoteDetailBiz(
			repo, requester)

		result, err := biz.SeeExportNoteDetail(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}

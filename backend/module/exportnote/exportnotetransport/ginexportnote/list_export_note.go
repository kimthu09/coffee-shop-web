package ginexportnote

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/exportnote/exportnotebiz"
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"coffee_shop_management_backend/module/exportnote/exportnoterepo"
	"coffee_shop_management_backend/module/exportnote/exportnotestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListExportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter exportnotemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := exportnotestore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := exportnoterepo.NewListExportNoteRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := exportnotebiz.NewListExportNoteRepo(repo, requester)

		result, err := biz.ListExportNote(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

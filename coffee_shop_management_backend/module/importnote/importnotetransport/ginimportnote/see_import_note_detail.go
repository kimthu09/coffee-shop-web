package ginimportnote

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnote/importnotebiz"
	"coffee_shop_management_backend/module/importnote/importnoterepo"
	"coffee_shop_management_backend/module/importnote/importnotestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeImportNoteDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		importNoteStore := importnotestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := importnoterepo.NewSeeImportNoteDetailRepo(importNoteStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := importnotebiz.NewSeeImportNoteDetailBiz(
			repo, requester)

		result, err := biz.SeeImportNoteDetail(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}

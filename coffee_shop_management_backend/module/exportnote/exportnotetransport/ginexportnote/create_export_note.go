package ginexportnote

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/exportnote/exportnotebiz"
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"coffee_shop_management_backend/module/exportnote/exportnotestore"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailstore"
	"coffee_shop_management_backend/module/ingredient/ingredientstore"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateExportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data exportnotemodel.ExportNoteCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		exportNoteStore := exportnotestore.NewSQLStore(appCtx.GetMainDBConnection())
		exportNoteDetailStore := exportnotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		ingredientStore := ingredientstore.NewSQLStore(appCtx.GetMainDBConnection())
		ingredientDetailStore := ingredientdetailstore.NewSQLStore(appCtx.GetMainDBConnection())

		requester := c.MustGet(common.CurrentUserStr).(common.Requester)
		data.CreateBy = requester.GetUserId()

		business := exportnotebiz.NewCreateExportNoteBiz(
			exportNoteStore,
			exportNoteDetailStore,
			ingredientStore,
			ingredientDetailStore,
		)

		if err := business.CreateExportNote(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.Id))
	}
}

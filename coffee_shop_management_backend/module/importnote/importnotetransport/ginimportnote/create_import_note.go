package ginimportnote

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/importnote/importnotebiz"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/importnote/importnotestore"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailstore"
	"coffee_shop_management_backend/module/ingredient/ingredientstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data importnotemodel.ImportNoteCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		importNoteStore := importnotestore.NewSQLStore(appCtx.GetMainDBConnection())
		importNoteDetailStore := importnotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		ingredientStore := ingredientstore.NewSQLStore(appCtx.GetMainDBConnection())

		requester := c.MustGet(common.CurrentUserStr).(common.Requester)
		data.CreateBy = requester.GetUserId()

		business := importnotebiz.NewCreateImportNoteBiz(
			importNoteStore,
			importNoteDetailStore,
			ingredientStore,
		)

		if err := business.CreateImportNote(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.Id))
	}
}

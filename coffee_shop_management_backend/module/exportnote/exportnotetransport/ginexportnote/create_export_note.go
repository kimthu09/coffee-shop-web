package ginexportnote

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/exportnote/exportnotebiz"
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"coffee_shop_management_backend/module/exportnote/exportnoterepo"
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

		requester := c.MustGet(common.CurrentUserStr).(common.Requester)
		data.CreateBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		exportNoteStore := exportnotestore.NewSQLStore(db)
		exportNoteDetailStore := exportnotedetailstore.NewSQLStore(db)
		ingredientStore := ingredientstore.NewSQLStore(db)
		ingredientDetailStore := ingredientdetailstore.NewSQLStore(db)

		repo := exportnoterepo.NewCreateCancelNoteRepo(
			exportNoteStore,
			exportNoteDetailStore,
			ingredientStore,
			ingredientDetailStore,
		)

		business := exportnotebiz.NewCreateExportNoteBiz(repo)

		if err := business.CreateExportNote(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.Id))
	}
}

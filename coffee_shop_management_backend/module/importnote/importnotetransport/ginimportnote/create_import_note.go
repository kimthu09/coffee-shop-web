package ginimportnote

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnote/importnotebiz"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/importnote/importnoterepo"
	"coffee_shop_management_backend/module/importnote/importnotestore"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailstore"
	"coffee_shop_management_backend/module/ingredient/ingredientstore"
	"coffee_shop_management_backend/module/supplier/supplierstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data importnotemodel.ImportNoteCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreateBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		importNoteStore := importnotestore.NewSQLStore(db)
		importNoteDetailStore := importnotedetailstore.NewSQLStore(db)
		ingredientStore := ingredientstore.NewSQLStore(db)
		supplierStore := supplierstore.NewSQLStore(db)

		repo := importnoterepo.NewCreateImportNoteRepo(
			importNoteStore,
			importNoteDetailStore,
			ingredientStore,
			supplierStore,
		)

		gen := generator.NewShortIdGenerator()

		business := importnotebiz.NewCreateImportNoteBiz(gen, repo, requester)

		if err := business.CreateImportNote(c.Request.Context(), &data); err != nil {
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

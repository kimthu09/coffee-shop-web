package ginimportnote

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/importnote/importnotebiz"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/importnote/importnotestore"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailstore"
	"coffee_shop_management_backend/module/ingredient/ingredientstore"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailstore"
	"coffee_shop_management_backend/module/supplier/supplierstore"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtstore"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeStatusImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idImportNote := c.Param("id")
		if idImportNote == "" {
			panic(common.ErrInvalidRequest(errors.New("param id not exist")))
		}

		var data importnotemodel.ImportNoteUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		importNoteStore := importnotestore.NewSQLStore(appCtx.GetMainDBConnection())
		importNoteDetailStore := importnotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		ingredientDetailStore := ingredientdetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		ingredientStore := ingredientstore.NewSQLStore(appCtx.GetMainDBConnection())
		supplierStore := supplierstore.NewSQLStore(appCtx.GetMainDBConnection())
		supplierDebtStore := supplierdebtstore.NewSQLStore(appCtx.GetMainDBConnection())

		requester := c.MustGet(common.CurrentUserStr).(common.Requester)
		data.CloseBy = requester.GetUserId()

		business := importnotebiz.NewChangeStatusImportNoteBiz(
			importNoteStore,
			importNoteDetailStore,
			ingredientStore,
			ingredientDetailStore,
			supplierStore,
			supplierDebtStore,
		)

		if err := business.ChangeStatusImportNote(
			c.Request.Context(),
			idImportNote, &data,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}

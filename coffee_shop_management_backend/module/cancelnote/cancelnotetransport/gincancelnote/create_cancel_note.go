package gincancelnote

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/cancelnote/cancelnotebiz"
	"coffee_shop_management_backend/module/cancelnote/cancelnotemodel"
	"coffee_shop_management_backend/module/cancelnote/cancelnotestore"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailstore"
	"coffee_shop_management_backend/module/ingredient/ingredientstore"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCancelNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data cancelnotemodel.CancelNoteCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		cancelNoteStore := cancelnotestore.NewSQLStore(appCtx.GetMainDBConnection())
		cancelNoteDetailStore := cancelnotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		ingredientStore := ingredientstore.NewSQLStore(appCtx.GetMainDBConnection())
		ingredientDetailStore := ingredientdetailstore.NewSQLStore(appCtx.GetMainDBConnection())

		requester := c.MustGet(common.CurrentUserStr).(common.Requester)
		data.CreateBy = requester.GetUserId()

		business := cancelnotebiz.NewCreateCancelNoteBiz(
			cancelNoteStore,
			cancelNoteDetailStore,
			ingredientStore,
			ingredientDetailStore,
		)

		if err := business.CreateCancelNote(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.Id))
	}
}

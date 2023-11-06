package gincancelnote

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/cancelnote/cancelnotebiz"
	"coffee_shop_management_backend/module/cancelnote/cancelnotemodel"
	"coffee_shop_management_backend/module/cancelnote/cancelnoterepo"
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

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreateBy = requester.GetUserId()

		gen := generator.NewShortIdGenerator()

		db := appCtx.GetMainDBConnection().Begin()

		cancelNoteStore := cancelnotestore.NewSQLStore(db)
		cancelNoteDetailStore := cancelnotedetailstore.NewSQLStore(db)
		ingredientStore := ingredientstore.NewSQLStore(db)
		ingredientDetailStore := ingredientdetailstore.NewSQLStore(db)

		repo := cancelnoterepo.NewCreateCancelNoteRepo(
			cancelNoteStore,
			cancelNoteDetailStore,
			ingredientStore,
			ingredientDetailStore,
		)

		business := cancelnotebiz.NewCreateCancelNoteBiz(gen, repo, requester)

		if err := business.CreateCancelNote(c.Request.Context(), &data); err != nil {
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

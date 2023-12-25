package gininventorychecknote

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/ingredient/ingredientstore"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotebiz"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotemodel"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknoterepo"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotestore"
	"coffee_shop_management_backend/module/inventorychecknotedetail/inventorychecknotedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateInventoryCheckNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data inventorychecknotemodel.InventoryCheckNoteCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreatedBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		inventoryCheckNoteStore := inventorychecknotestore.NewSQLStore(db)
		inventoryCheckNoteDetailStore := inventorychecknotedetailstore.NewSQLStore(db)
		ingredientStore := ingredientstore.NewSQLStore(db)

		repo := inventorychecknoterepo.NewCreateInventoryCheckNoteRepo(
			inventoryCheckNoteStore,
			inventoryCheckNoteDetailStore,
			ingredientStore,
		)

		gen := generator.NewShortIdGenerator()

		business := inventorychecknotebiz.NewCreateInventoryCheckNoteBiz(gen, repo, requester)

		if err := business.CreateInventoryCheckNote(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}

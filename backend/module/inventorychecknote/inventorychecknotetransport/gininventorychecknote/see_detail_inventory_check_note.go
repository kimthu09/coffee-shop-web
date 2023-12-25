package gininventorychecknote

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotebiz"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknoterepo"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotestore"
	"coffee_shop_management_backend/module/inventorychecknotedetail/inventorychecknotedetailstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeDetailInventoryCheckNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		inventoryCheckNoteId := c.Param("id")

		inventoryCheckNoteStore :=
			inventorychecknotestore.NewSQLStore(appCtx.GetMainDBConnection())
		inventoryCheckNoteDetailStore :=
			inventorychecknotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := inventorychecknoterepo.NewSeeDetailInventoryCheckNoteRepo(
			inventoryCheckNoteStore, inventoryCheckNoteDetailStore)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := inventorychecknotebiz.NewSeeDetailImportNoteBiz(repo, requester)

		result, err := biz.SeeDetailInventoryCheckNote(
			c.Request.Context(), inventoryCheckNoteId)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}

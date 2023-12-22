package gininventorychecknote

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotebiz"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotemodel"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknoterepo"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListInventoryCheckNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter inventorychecknotemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := inventorychecknotestore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := inventorychecknoterepo.NewListInventoryCheckNoteRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := inventorychecknotebiz.NewListInventoryCheckNoteBiz(repo, requester)

		result, err := biz.ListInventoryCheckNote(
			c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

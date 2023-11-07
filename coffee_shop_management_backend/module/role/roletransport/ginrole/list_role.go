package ginrole

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/role/rolebiz"
	"coffee_shop_management_backend/module/role/rolerepo"
	"coffee_shop_management_backend/module/role/rolestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRole(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := rolestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := rolerepo.NewListRoleRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := rolebiz.NewListRoleBiz(repo, requester)

		result, err := biz.ListRole(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))
	}
}

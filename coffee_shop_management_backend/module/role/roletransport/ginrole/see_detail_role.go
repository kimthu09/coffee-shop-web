package ginrole

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/feature/featurestore"
	"coffee_shop_management_backend/module/role/rolebiz"
	"coffee_shop_management_backend/module/role/rolerepo"
	"coffee_shop_management_backend/module/role/rolestore"
	"coffee_shop_management_backend/module/rolefeature/rolefeaturestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeeDetailRole(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		db := appCtx.GetMainDBConnection()
		roleStore := rolestore.NewSQLStore(db)
		roleFeatureStore := rolefeaturestore.NewSQLStore(db)
		featureStore := featurestore.NewSQLStore(db)

		repo := rolerepo.NewSeeRoleDetailRepo(roleStore, roleFeatureStore, featureStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := rolebiz.NewSeeDetailRoleBiz(repo, requester)

		result, err := biz.SeeDetailRole(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}

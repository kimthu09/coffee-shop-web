package ginrole

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/role/rolebiz"
	"coffee_shop_management_backend/module/role/rolemodel"
	"coffee_shop_management_backend/module/role/rolerepo"
	"coffee_shop_management_backend/module/role/rolestore"
	"coffee_shop_management_backend/module/rolefeature/rolefeaturestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateRole(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data rolemodel.RoleUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		roleStore := rolestore.NewSQLStore(db)
		roleFeatureStore := rolefeaturestore.NewSQLStore(db)

		repo := rolerepo.NewUpdateRoleRepo(
			roleStore,
			roleFeatureStore,
		)

		business := rolebiz.NewUpdateRoleBiz(repo, requester)

		if err := business.UpdateRole(c.Request.Context(), id, &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

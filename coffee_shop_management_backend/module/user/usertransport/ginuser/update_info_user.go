package ginuser

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/user/userbiz"
	"coffee_shop_management_backend/module/user/usermodel"
	"coffee_shop_management_backend/module/user/userrepo"
	"coffee_shop_management_backend/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateInfoUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data usermodel.UserUpdateInfo

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := userstore.NewSQLStore(db)
		repo := userrepo.NewUpdateInfoUserRepo(store)

		business := userbiz.NewUpdateInfoUserBiz(repo, requester)

		if err := business.UpdateUser(c.Request.Context(), id, &data); err != nil {
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

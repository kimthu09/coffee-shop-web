package ginuser

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/module/user/userbiz"
	"coffee_shop_management_backend/module/user/userrepo"
	"coffee_shop_management_backend/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := userstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := userrepo.NewGetAllUserRepo(store)

		biz := userbiz.NewGetAllUserBiz(repo)

		result, err := biz.GetAllUser(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}

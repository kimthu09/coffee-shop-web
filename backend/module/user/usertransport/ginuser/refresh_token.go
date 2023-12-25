package ginuser

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/appctx"
	"coffee_shop_management_backend/component/tokenprovider/jwt"
	"coffee_shop_management_backend/module/user/userbiz"
	"coffee_shop_management_backend/module/user/usermodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshToken(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserRefreshToken
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

		business := userbiz.NewRefreshTokenBiz(60*60*24, tokenProvider)
		account, err := business.RefreshToken(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}

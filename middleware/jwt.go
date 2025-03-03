package middleware

import (
	"errors"
	"net/http"

	"github.com/goravel/framework/auth"
	contractshttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Jwt() contractshttp.Middleware {
	return func(ctx contractshttp.Context) {
		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, &contractshttp.Json{
				"code":    http.StatusUnauthorized,
				"message": "token缺失",
			})
			return
		}

		if _, err := facades.Auth(ctx).Parse(token); err != nil {
			if errors.Is(err, auth.ErrorTokenExpired) {
				token, err = facades.Auth(ctx).Refresh()
				if err != nil {
					ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, &contractshttp.Json{
						"code":    http.StatusUnauthorized,
						"message": "刷新token过期",
					})
					return
				}
				token = "Bearer " + token
			} else {

				ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, &contractshttp.Json{
					"code":    http.StatusUnauthorized,
					"message": "token过期",
				})
				return
			}
		}

		// You can get User in DB and set it to ctx

		//var user models.User
		//if err := facades.Auth().User(ctx, &user); err != nil {
		//	ctx.Request().AbortWithStatus(http.StatusUnauthorized)
		//  return
		//}
		//ctx.WithValue("user", user)

		ctx.Response().Header("Authorization", token)
		ctx.Request().Next()
	}
}

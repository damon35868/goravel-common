package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/goravel/framework/auth"
	contractshttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Jwt(messageMaps ...map[string]string) contractshttp.Middleware {
	var (
		missingTokenMsg = "未携带token"
		parseMsg        = "传入了非法的token内容，解析失败"
		expiredMsg      = "token过期，请重新登录～"
	)
	if len(messageMaps) > 0 {
		msgMap := messageMaps[0]
		if msgMap["missingTokenMsg"] != "" {
			missingTokenMsg = msgMap["missingTokenMsg"]
		}
		if msgMap["parseMsg"] != "" {
			parseMsg = msgMap["parseMsg"]
		}
		if msgMap["expiredMsg"] != "" {
			expiredMsg = msgMap["expiredMsg"]
		}
	}

	return func(ctx contractshttp.Context) {
		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, &contractshttp.Json{
				"code":    http.StatusUnauthorized,
				"message": missingTokenMsg,
			})
			return
		}

		payload, err := facades.Auth(ctx).Parse(token)
		if err != nil {
			if errors.Is(err, auth.ErrorTokenExpired) {
				ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, &contractshttp.Json{
					"code":    http.StatusUnauthorized,
					"message": expiredMsg,
				})
				return
			}
			if errors.Is(err, auth.ErrorInvalidToken) {
				ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, &contractshttp.Json{
					"code":    http.StatusUnauthorized,
					"message": parseMsg,
				})
				return
			}
		}
		if payload == nil || payload.Key == "" {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, &contractshttp.Json{
				"code":    http.StatusUnauthorized,
				"message": parseMsg,
			})
			return
		}

		ctx.Request().Next()
	}
}

func JwtWithSSO(messageMaps ...map[string]string) contractshttp.Middleware {
	var (
		missingTokenMsg = "未携带token"
		ssoMsg          = "当前账号已在其他地方登录，请重新登录～"
		parseMsg        = "传入了非法的token内容，解析失败"
		expiredMsg      = "token过期，请重新登录～"
	)
	if len(messageMaps) > 0 {
		msgMap := messageMaps[0]
		if msgMap["missingTokenMsg"] != "" {
			missingTokenMsg = msgMap["missingTokenMsg"]
		}
		if msgMap["ssoMsg"] != "" {
			ssoMsg = msgMap["ssoMsg"]
		}
		if msgMap["parseMsg"] != "" {
			parseMsg = msgMap["parseMsg"]
		}
		if msgMap["expiredMsg"] != "" {
			expiredMsg = msgMap["expiredMsg"]
		}
	}

	return func(ctx contractshttp.Context) {
		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, &contractshttp.Json{
				"code":    http.StatusUnauthorized,
				"message": missingTokenMsg,
			})
			return
		}

		payload, err := facades.Auth(ctx).Parse(token)
		if err != nil {
			if errors.Is(err, auth.ErrorTokenExpired) {
				ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, &contractshttp.Json{
					"code":    http.StatusUnauthorized,
					"message": expiredMsg,
				})
				return
			}
			if errors.Is(err, auth.ErrorInvalidToken) {
				ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, &contractshttp.Json{
					"code":    http.StatusUnauthorized,
					"message": parseMsg,
				})
				return
			}
		}
		if payload == nil || payload.Key == "" {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, &contractshttp.Json{
				"code":    http.StatusUnauthorized,
				"message": parseMsg,
			})
			return
		}

		config := facades.Config()
		sso := config.GetBool("jwt.sso")

		if sso {
			cacheToken := facades.Cache().Get(fmt.Sprintf("jwt:user:%s", payload.Key))
			if cacheToken != nil && cacheToken.(string) != strings.TrimPrefix(token, "Bearer ") {
				ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, &contractshttp.Json{
					"code":    http.StatusUnauthorized,
					"message": ssoMsg,
				})
				return
			}
		}

		ctx.Request().Next()
	}
}

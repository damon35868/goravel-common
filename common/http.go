package common

import (
	"goravel-common/errorx"
	"strconv"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type ResponseOptions struct {
	Code    int
	Message *string
}

/**
 * @description: 公共统一返回结构
 * @param {http.Context} ctx
 * @param {interface{}} data
 * @param {...ResponseOptions} options
 * @return {*}
 */
func Response(ctx http.Context, data interface{}, options ...ResponseOptions) http.Response {
	code, message := http.StatusOK, "请求成功"

	if len(options) > 0 {
		if options[0].Code != 0 {
			code = options[0].Code
		}
		if options[0].Message != nil {
			message = *options[0].Message
		}
	}

	return ctx.Response().Success().Json(&http.Json{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

/**
 * @description: 分页: 是否有下一页判断
 * @param {*} skip
 * @param {*} take
 * @param {int64} total
 * @return {*}
 */
func HasNextPage(skip, take, total int64) bool {
	return skip*take < total
}

/**
 * @description: 获取token user ID
 * @param {http.Context} ctx
 * @return {*}
 */
func GetTokenUserId(ctx http.Context) int64 {
	idStr, _ := facades.Auth(ctx).ID()
	id, _ := strconv.Atoi(idStr)

	return int64(id)
}

/**
 * @description: 校验request并输出错误
 * @param {http.Context} ctx
 * @param {http.FormRequest} val
 * @return {*}
 */
func ValidateRequest(ctx http.Context, val http.FormRequest) http.Response {
	errors, err := ctx.Request().ValidateRequest(val)

	if err != nil {
		return errorx.HttpException(ctx, http.StatusBadRequest, err.Error())
	}
	if errors != nil {
		return errorx.HttpExceptionAndData(ctx, http.StatusBadRequest, errors.All())
	}
	return nil
}

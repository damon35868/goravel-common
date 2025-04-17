package errorx

import "github.com/goravel/framework/contracts/http"

/**
 * @description: http异常
 * @param {http.Context} ctx
 * @param {int} errCode
 * @param {...string} msgs
 * @return {*}
 */
func HttpException(ctx http.Context, errCode int, msgs ...string) http.Response {
	code, msg := http.StatusInternalServerError, MapErrMsg(errCode, "未知错误")
	if errCode != 0 {
		code = errCode
	}
	if len(msgs) > 0 {
		msg = msgs[0]
	}
	return ctx.Response().Status(code).Json(&http.Json{
		"code":    code,
		"message": msg,
	})
}

/**
 * @description: http异常 - 中断
 * @param {http.Context} ctx
 * @param {int} errCode
 * @param {...string} msgs
 * @return {*}
 */
func HttpExceptionAbortable(ctx http.Context, err any, errCodes ...int) error {
	code, err := http.StatusInternalServerError, err
	if len(errCodes) > 0 {
		code = errCodes[0]
	}
	return ctx.Response().Status(code).Json(&http.Json{
		"code":    code,
		"message": err,
	}).Abort()
}

/**
 * @description: http异常并输出数据
 * @param {http.Context} ctx
 * @param {int} errCode
 * @param {interface{}} data
 * @param {...string} msgs
 * @return {*}
 */
func HttpExceptionAndData(ctx http.Context, errCode int, data interface{}, msgs ...string) http.Response {
	code, msg := http.StatusInternalServerError, MapErrMsg(errCode, "未知错误")
	if errCode != 0 {
		code = errCode
	}
	if len(msgs) > 0 {
		msg = msgs[0]
	}
	return ctx.Response().Status(code).Json(&http.Json{
		"code":    code,
		"message": msg,
		"data":    data,
	})
}

/**
 * @description: 业务异常
 * @param {http.Context} ctx
 * @param {int} errCode
 * @param {...string} msgs
 * @return {*}
 */
func BusinessException(ctx http.Context, errCode int, msgs ...string) http.Response {
	code, msg := BASE_ERROR, MapErrMsg(errCode)
	if errCode != 0 {
		code = errCode
	}
	if len(msgs) > 0 {
		msg = msgs[0]
	}
	return ctx.Response().Success().Json(&http.Json{
		"code":    code,
		"message": msg,
	})

}

package errorx

var (
	BASE_ERROR = 100000
	MessageMap = map[int]string{
		BASE_ERROR: "当前业务繁忙，请稍后重试~",
	}
)

func Boot(errorCodes map[int]string) {
	for code, msg := range errorCodes {
		MessageMap[code] = msg
	}
}

func MapErrMsg(errCode int, defaultMsgs ...string) string {
	if msg, ok := MessageMap[errCode]; ok {
		return msg
	} else {
		if len(defaultMsgs) > 0 {
			return defaultMsgs[0]
		}
		return MessageMap[BASE_ERROR]
	}
}

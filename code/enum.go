/**
 * @Author: Resynz
 * @Date: 2021/7/19 13:42
 */
package code

type ResponseCode int

const (
	SuccessCode ResponseCode = 0

	InvalidAuth ResponseCode = 401

	// general error code
	BadRequest     ResponseCode = 1000
	InvalidRequest ResponseCode = 1000 + iota
	InvalidParams
)

var ResponseCodeMap = map[ResponseCode]string{
	SuccessCode:    "请求成功",
	InvalidAuth:    "身份校验失败",
	InvalidRequest: "无效的请求",
	BadRequest:     "系统错误",
	InvalidParams:  "请求参数无效",
}

func GetCodeMsg(code ResponseCode) string {
	h, ok := ResponseCodeMap[code]
	if !ok {
		return "未知错误"
	}
	return h
}

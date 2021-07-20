/**
 * @Author: Resynz
 * @Date: 2021/7/20 10:36
 */
package api

import (
	"ws-service/code"
	"ws-service/common"
	"ws-service/strategy"
)

// GetWsUrl 获取WebSocketUrl
func GetWsUrl(ctx *common.Context) {
	type formValidate struct {
		UserId string `form:"user_id" binding:"required" json:"user_id"`
	}
	var form formValidate
	if err := ctx.ShouldBind(&form); err != nil {
		common.HandleResponse(ctx, code.InvalidParams, nil)
		return
	}
	wsUrl := ""
	wsMap := strategy.GetWsServer(form.UserId)
	if wsMap != nil {
		wsUrl = wsMap.WsUrl
	}
	data := map[string]string{
		"ws_url": wsUrl,
	}
	common.HandleResponse(ctx, code.SuccessCode, data)
}

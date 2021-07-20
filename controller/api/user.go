/**
 * @Author: Resynz
 * @Date: 2021/7/20 13:46
 */
package api

import (
	"ws-service/code"
	"ws-service/common"
	"ws-service/config"
	"ws-service/lib/ws"
)

// GetOnlineCount 获取在线人数
func GetOnlineCount(ctx *common.Context) {
	var count int64
	for _, v := range config.Conf.WsList {
		c, err := ws.OnlineCount(v)
		if err != nil {
			ctx.Logger.Errorf("[GetOnlineCount] error:%s", err)
			common.HandleResponse(ctx, code.BadRequest, nil, err.Error())
			return
		}
		count += c
	}
	data := map[string]int64{
		"count": count,
	}
	common.HandleResponse(ctx, code.SuccessCode, data)
}

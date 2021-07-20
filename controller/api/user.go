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
	"ws-service/strategy"
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

// IsOnline 是否在线
func IsOnline(ctx *common.Context) {
	type formValidate struct {
		UserId string `form:"user_id" binding:"required" json:"user_id"`
	}
	var form formValidate
	if err := ctx.ShouldBind(&form); err != nil {
		common.HandleResponse(ctx, code.InvalidParams, nil, err.Error())
		return
	}
	wp := strategy.GetWsServer(form.UserId)
	if wp == nil {
		data := map[string]bool{
			"result": false,
		}
		common.HandleResponse(ctx, code.SuccessCode, data)
		return
	}
	result, err := ws.IsOnline(wp, form.UserId)
	if err != nil {
		common.HandleResponse(ctx, code.BadRequest, nil, err.Error())
		return
	}
	data := map[string]bool{
		"result": result,
	}
	common.HandleResponse(ctx, code.SuccessCode, data)
}

// UserInfo
func UserInfo(ctx *common.Context) {
	type formValidate struct {
		UserId string `form:"user_id" binding:"required" json:"user_id"`
	}
	var form formValidate
	if err := ctx.ShouldBind(&form); err != nil {
		common.HandleResponse(ctx, code.InvalidParams, nil, err.Error())
		return
	}
	wp := strategy.GetWsServer(form.UserId)
	if wp == nil {
		common.HandleResponse(ctx, code.InvalidRequest, nil, "ws server not config")
		return
	}
	info, err := ws.UserInfo(wp, form.UserId)
	if err != nil {
		common.HandleResponse(ctx, code.BadRequest, nil, err.Error())
		return
	}
	data := map[string]interface{}{
		"info": info,
	}
	common.HandleResponse(ctx, code.SuccessCode, data)
}

/**
 * @Author: Resynz
 * @Date: 2021/7/20 15:09
 */
package api

import (
	"ws-service/code"
	"ws-service/common"
	"ws-service/config"
	"ws-service/lib/ws"
	"ws-service/util"
)

// SendMsg 发送信息
func SendMsg(ctx *common.Context) {
	type formValidate struct {
		MsgList      []string `form:"msg_list" binding:"required" json:"msg_list"`
		UserIdList   []string `form:"user_id_list" binding:"required" json:"user_id_list"`
		ClientIdList []string `form:"client_id_list" binding:"" json:"client_id_list"`
	}
	var form formValidate
	if err := ctx.ShouldBind(&form); err != nil {
		common.HandleResponse(ctx, code.InvalidParams, nil, err.Error())
		return
	}

	if len(config.Conf.WsList) == 0 {
		common.HandleResponse(ctx, code.InvalidRequest, nil, "ws server not config")
		return
	}

	wsMap := make(map[int][]string)
	for _, c := range form.UserIdList {
		number := util.ComputedNumberByString(c)
		index := number % len(config.Conf.WsList)
		var cs []string
		if v, ok := wsMap[index]; ok {
			cs = v
			cs = append(cs, c)
		} else {
			cs = make([]string, 1)
			cs[0] = c
		}
		wsMap[index] = cs
	}

	for k, v := range wsMap {
		if err := ws.SendMsg(config.Conf.WsList[k], v, form.MsgList, form.ClientIdList); err != nil {
			common.HandleResponse(ctx, code.BadRequest, nil, err.Error())
			return
		}
	}

	data := map[string]bool{
		"result": true,
	}

	common.HandleResponse(ctx, code.SuccessCode, data)
}

// Broadcast 广播
func Broadcast(ctx *common.Context) {
	type formValidate struct {
		MsgList []string `form:"msg_list" binding:"required" json:"msg_list"`
	}
	var form formValidate
	if err := ctx.ShouldBind(&form); err != nil {
		common.HandleResponse(ctx, code.InvalidParams, nil)
		return
	}
	if len(config.Conf.WsList) == 0 {
		common.HandleResponse(ctx, code.InvalidRequest, nil, "ws server not config")
		return
	}
	for _, v := range config.Conf.WsList {
		if err := ws.Broadcast(v, form.MsgList); err != nil {
			common.HandleResponse(ctx, code.BadRequest, nil, err.Error())
			return
		}
	}
	data := map[string]bool{
		"result": true,
	}

	common.HandleResponse(ctx, code.SuccessCode, data)
}

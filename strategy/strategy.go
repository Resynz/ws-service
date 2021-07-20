package strategy

import (
	"ws-service/config"
	"ws-service/util"
)

// 分配用户到ws-server策略
func GetWsServer(client string) *config.WsMap {
	if len(config.Conf.WsList) == 0 {
		return nil
	}
	number := util.ComputedNumberByString(client)
	index := number % len(config.Conf.WsList)
	return config.Conf.WsList[index]
}

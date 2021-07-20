/**
 * @Author: Resynz
 * @Date: 2021/7/19 18:03
 */
package config

type logConf struct {
	Path         string `json:"path"`
	Name         string `json:"name"`
	Level        string `json:"level"`
	RequestIdKey string `json:"request_id_key"`
}

type WxMap struct {
	WsUrl   string `json:"ws_url"`
	BaseUrl string `json:"base_url"`
}

type Config struct {
	Mode      string   `json:"mode"`
	AppPort   int      `json:"app_port"`
	LogConfig logConf  `json:"log_config"`
	WsList    []*WxMap `json:"ws_list"`
}

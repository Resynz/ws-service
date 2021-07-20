/**
 * @Author: Resynz
 * @Date: 2021/7/20 13:47
 */
package ws

import (
	"encoding/json"
	"fmt"
	"github.com/rosbit/go-wget"
	"net/http"
	"ws-service/config"
)

type BaseRes struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
}

func OnlineCount(w *config.WxMap) (int64, error) {
	reqUrl := fmt.Sprintf("%s/api/online-count", w.BaseUrl)
	method := "GET"
	status, content, _, err := wget.Wget(reqUrl, method, nil, nil)
	if err != nil {
		return 0, err
	}
	if status != http.StatusOK {
		return 0, fmt.Errorf("bad http status:%d", status)
	}

	type res struct {
		*BaseRes
		Data struct {
			Count int64 `json:"count"`
		} `json:"data"`
	}

	var r res
	if err = json.Unmarshal(content, &r); err != nil {
		return 0, err
	}

	return r.Data.Count, nil
}

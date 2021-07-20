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

func OnlineCount(w *config.WsMap) (int64, error) {
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

func SendMsg(w *config.WsMap, userIds, messages []string) error {
	reqUrl := fmt.Sprintf("%s/api/send-msg", w.BaseUrl)
	method := "POST"
	params := map[string]interface{}{
		"user_id_list": userIds,
		"msg_list":     messages,
	}
	status, _, _, err := wget.PostJson(reqUrl, method, params, nil)
	if err != nil {
		return err
	}
	if status != http.StatusOK {
		return fmt.Errorf("bad http status:%d", status)
	}
	return nil
}

func Broadcast(w *config.WsMap, messages []string) error {
	reqUrl := fmt.Sprintf("%s/api/broadcast", w.BaseUrl)
	method := "POST"
	params := map[string]interface{}{
		"msg_list": messages,
	}
	status, _, _, err := wget.PostJson(reqUrl, method, params, nil)
	if err != nil {
		return err
	}
	if status != http.StatusOK {
		return fmt.Errorf("bad http status:%d", status)
	}
	return nil
}

func IsOnline(w *config.WsMap, userId string) (bool, error) {
	reqUrl := fmt.Sprintf("%s/api/is-online?user_id=%s", w.BaseUrl, userId)
	method := "GET"
	status, content, _, err := wget.Wget(reqUrl, method, nil, nil)
	if err != nil {
		return false, err
	}
	if status != http.StatusOK {
		return false, fmt.Errorf("bad http status:%d", status)
	}

	type res struct {
		*BaseRes
		Data struct {
			Result bool `json:"result"`
		} `json:"data"`
	}

	var r res
	if err = json.Unmarshal(content, &r); err != nil {
		return false, err
	}

	return r.Data.Result, nil
}

type ClientObj struct {
	ClientId   string `json:"client_id"`
	CreateTime int64  `json:"create_time"`
	Platform   int    `json:"platform"`
}

type UserInfoObj struct {
	UserId  string       `json:"user_id"`
	Clients []*ClientObj `json:"clients"`
}

type UserInfoResponse struct {
	*BaseRes
	Data struct {
		Info *UserInfoObj `json:"info"`
	} `json:"data"`
}

func UserInfo(w *config.WsMap, userId string) (*UserInfoObj, error) {
	reqUrl := fmt.Sprintf("%s/api/info?user_id=%s", w.BaseUrl, userId)
	method := "GET"
	status, content, _, err := wget.Wget(reqUrl, method, nil, nil)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("bad http status:%d", status)
	}

	var r UserInfoResponse
	if err = json.Unmarshal(content, &r); err != nil {
		return nil, err
	}

	return r.Data.Info, nil
}

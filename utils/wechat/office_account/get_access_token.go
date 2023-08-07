package office_account

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// GetAccessToken 微信公众号平台获取Token 返回AccessToken 和过期时间 source:https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func GetAccessToken(appID, appSecret string) (*AccessTokenResp, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appID, appSecret)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON响应
	var result AccessTokenResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ExpiresIn == 0 {
		return nil, errors.New("get access token error")
	}
	return &result, nil
}

type AccessTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

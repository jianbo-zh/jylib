package wxutil

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GetOpenId 获取微信OpendID
func GetOpenId(ctx context.Context, appId string, appSecret string, jsCode string) (string, error) {
	params := url.Values{}
	params.Add("appid", appId)
	params.Add("secret", appSecret)
	params.Add("grant_type", "authorization_code")
	params.Add("js_code", jsCode)

	requestURL := fmt.Sprintf("%s?%s", "https://api.weixin.qq.com/sns/jscode2session", params.Encode())

	// 发送HTTP GET请求
	resp, err := http.Get(requestURL)
	if err != nil {
		return "", fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("io.ReadAll error: %w", err)
	}

	// 解析JSON响应
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("json.Unmarshal error: %w", err)
	}

	var openid string
	if oid, ok := result["openid"]; !ok {
		return "", fmt.Errorf("no openid, body: %s", body)

	} else if openid, ok = oid.(string); !ok {
		return "", fmt.Errorf("openid type assert failed")
	}

	return openid, nil
}

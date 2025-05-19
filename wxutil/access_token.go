package wxutil

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetAccessToken(ctx context.Context, appId string, appSecret string) (*AccessToken, error) {
	u, _ := url.Parse("https://api.weixin.qq.com/cgi-bin/token")
	query := u.Query()
	query.Add("grant_type", "client_credential")
	query.Add("appid", appId)
	query.Add("secret", appSecret)
	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var wxAT AccessToken
	if err := json.Unmarshal(bs, &wxAT); err != nil {
		return nil, fmt.Errorf("json.Unmarshall error: %w", err)
	}

	return &wxAT, nil
}

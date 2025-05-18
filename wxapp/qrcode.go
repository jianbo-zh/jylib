package wxapp

import (
	"fmt"
	"net/url"
	"strings"
)

func QrcodeUrl(appid string, path string, query map[string]string) string {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "open.weixin.qq.com"
	u.Path = "/sns/getexpappinfo"
	u.Fragment = "wechat-redirect"

	values := u.Query()
	values.Add("appid", appid)

	if len(path) > 0 {
		if len(query) > 0 {
			var querys []string
			for k, v := range query {
				querys = append(querys, fmt.Sprintf("%s=%s", k, v))
			}

			path += "?" + strings.Join(querys, "&")
		}

		values.Add("path", path)
	}

	u.RawQuery = values.Encode()

	return u.String()
}

func QrcodeUser(apiDomain string, query map[string]string) string {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = apiDomain
	u.Path = "/qrcode/user"

	if len(query) > 0 {
		values := u.Query()
		for k, v := range query {
			values.Add(k, v)
		}

		u.RawQuery = values.Encode()
	}

	return u.String()
}

func QrcodeAdmin(apiDomain string, query map[string]string) string {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = apiDomain
	u.Path = "/qrcode/admin"

	if len(query) > 0 {
		values := u.Query()
		for k, v := range query {
			values.Add(k, v)
		}

		u.RawQuery = values.Encode()
	}

	return u.String()
}

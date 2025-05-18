package jwt

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type anyClaims struct {
	Payload []byte `json:"payload,omitempty"`
	jwt.RegisteredClaims
}

// Encode 生成JWT token
func Encode[T any](data T, key []byte, expiredAt time.Time) (string, error) {

	bs, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("json.marshal error: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, anyClaims{
		Payload: bs,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
	})

	return token.SignedString(key)
}

// Decode 解析JWT token
func Decode[T any](token string, key []byte) (data T, err error) {
	var t T

	jwtToken, err := jwt.ParseWithClaims(token, &anyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
	if err != nil {
		return t, fmt.Errorf("jwt.ParseWithClaims error: %w", err)
	}

	claims, ok := jwtToken.Claims.(*anyClaims)
	if !ok {
		return t, fmt.Errorf("jwtToken.Claims type assert error")
	}

	err = json.Unmarshal(claims.Payload, &data)
	if err != nil {
		return t, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	return data, nil
}

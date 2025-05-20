package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func PasswordHash(password string, account string) string {
	salt := "YUYHGr7@GF546DER6#@!!"
	data := fmt.Sprintf("%s_%d_%s_%s_%s", "uuid", 999993, salt, password, account)
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

func PasswordHashOps(password string) string {
	salt := "YUYHGr7@GF546DER6#@!!"
	data := fmt.Sprintf("%s_%d_%s_%s", "uuid", 999993, salt, password)
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

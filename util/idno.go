package util

import (
	"strings"
	"time"

	"github.com/jianbo-zh/jylib/typec"
)

func GenIdNo(typ typec.IDNo) string {
	return string(typ) + strings.Replace(time.Now().Format("20060102150405.000"), ".", "", 1)
}

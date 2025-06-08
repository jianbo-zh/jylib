package util

import (
	"fmt"
	"time"
)

func FormatFileSize(bytes int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	if bytes == 0 {
		return "0 B"
	}

	// 通过数学计算直接确定单位索引
	unitIndex := 0
	value := float64(bytes)
	for ; value >= 1024 && unitIndex < len(units)-1; unitIndex++ {
		value /= 1024
	}

	// 格式化输出
	if unitIndex == 0 {
		return fmt.Sprintf("%d B", bytes)
	} else {
		return fmt.Sprintf("%.2f %s", value, units[unitIndex])
	}
}

func FormatTime(ts time.Time, typeValue int) string {
	if ts.IsZero() {
		return ""
	}
	switch typeValue {
	case 0:
		return ts.Local().Format("2006-01-02 15:04:05")
	case 1:
		return ts.Local().Format("2006-01-02")
	case 2:
		return ts.Local().Format("15:04:05")
	default:
		return ""
	}
}

package util

import "fmt"

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

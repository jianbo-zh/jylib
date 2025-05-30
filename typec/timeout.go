package typec

import "time"

const (
	TaskTimeout_OrderAbortCheck          = 1 * time.Minute  // 1分钟
	TaskTimeout_OrderCancelCheck         = 5 * time.Minute  // 5分钟
	TaskTimeout_OrderExpirationCheck     = 24 * time.Hour   // 24小时
	TaskTimeout_OrderRefund              = 1 * time.Second  // 1秒
	TaskTimeout_OrderProfitSharingResult = 10 * time.Minute // 10分钟
)

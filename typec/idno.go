package typec

// ID编号类型（订单/押金/退款/分润）
type IDNo string

const (
	IDNo_Order   IDNo = "JY" // 订单编号
	IDNo_Deposit IDNo = "DN" // 订单押金编号
	IDNo_Refund  IDNo = "RN" // 订单退款编号
	IDNo_Sharing IDNo = "SN" // 订单分润编号
	IDNo_Flight  IDNo = "FL" // 班次任务编号
	IDNo_Coupon  IDNo = "CP" // 优惠卷编号
)

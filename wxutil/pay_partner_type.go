package wxutil

/*
	示例数据

	{
		"id": "EV-2018022511223320873",
		"create_time": "2015-05-20T13:29:35+08:00",
		"resource_type": "encrypt-resource",
		"event_type": "TRANSACTION.SUCCESS",
		"summary": "支付成功",
		"resource": {
			"original_type": "transaction",
			"algorithm": "AEAD_AES_256_GCM",
			"ciphertext": "",
			"associated_data": "",
			"nonce": ""
		}
	}

	{
		"sp_appid" : "wx8888888888888888",
		"sp_mchid" : "1230000109",
		"sub_appid" : "wxd678efh567hg6999",
		"sub_mchid" : "1900000109",
		"out_trade_no" : "1217752501201407033233368018",
		"transaction_id" : "1217752501201407033233368018",
		"trade_type" : "JSAPI",
		"trade_state" : "SUCCESS",
		"trade_state_desc" : "支付失败，请重新下单支付",
		"bank_type" : "CMC",
		"attach" : "自定义数据",
		"success_time" : "2018-06-08T10:34:56+08:00",
		"payer" : {
			"sp_openid" : "oUpF8uMuAJO_M2pxb1Q9zNjWeS6o\t",
			"sub_openid" : "oUpF8uMuAJO_M2pxb1Q9zNjWeS6o\t"
		},
		"amount" : {
			"total" : 100,
			"payer_total" : 100,
			"currency" : "CNY",
			"payer_currency" : "CNY"
		},
		"scene_info" : {
			"device_id" : "013467007045764"
		},
		"promotion_detail" : [
			{
			"coupon_id" : "109519",
			"name" : "单品惠-6",
			"scope" : "GLOBAL",
			"type" : "CASH",
			"amount" : 100,
			"stock_id" : "931386",
			"wechatpay_contribute" : 0,
			"merchant_contribute" : 0,
			"other_contribute" : 0,
			"currency" : "CNY",
			"goods_detail" : [
				{
				"goods_id" : "M1006",
				"quantity" : 1,
				"unit_price" : 100,
				"discount_amount" : 1,
				"goods_remark" : "商品备注信息"
				}
			]
			}
		]
	}
*/

type WxPayData struct {
	SpMchid        string `json:"sp_mchid,omitempty"`
	SpAppid        string `json:"sp_appid,omitempty"`
	SubMchid       string `json:"sub_mchid,omitempty"`
	SubAppid       string `json:"sub_appid,omitempty"`
	OutTradeNo     string `json:"out_trade_no,omitempty"`
	TransactionId  string `json:"transaction_id,omitempty"`
	TradeType      string `json:"trade_type,omitempty"`
	TradeState     string `json:"trade_state,omitempty"`
	TradeStateDesc string `json:"trade_state_desc,omitempty"`
	BankType       string `json:"bank_type,omitempty"`
	Attach         string `json:"attach,omitempty"`
	SuccessTime    string `json:"success_time,omitempty"`
	Payer          struct {
		SpOpenid  string `json:"sp_openid,omitempty"`
		SubOpenid string `json:"sub_openid,omitempty"`
	} `json:"payer,omitempty"`
	Amount struct {
		Total         int64  `json:"total,omitempty"`
		PayerTotal    int64  `json:"payer_total,omitempty"`
		Currency      string `json:"currency,omitempty"`
		PayerCurrency string `json:"payer_currency,omitempty"`
	} `json:"amount,omitempty"`
	// SceneInfo struct {
	// 	DeviceID string `json:"device_id,omitempty"`
	// } `json:"scene_info,omitempty"`
	// PromotionDetail []struct {
	// 	CouponID            string `json:"coupon_id,omitempty"`
	// 	Name                string `json:"name,omitempty"`
	// 	Scope               string `json:"scope,omitempty"`
	// 	Type                string `json:"type,omitempty"`
	// 	Amount              int64  `json:"amount,omitempty"`
	// 	StockID             string `json:"stock_id,omitempty"`
	// 	WechatpayContribute int64  `json:"wechatpay_contribute,omitempty"`
	// 	MerchantContribute  int64  `json:"merchant_contribute,omitempty"`
	// 	OtherContribute     int64  `json:"other_contribute,omitempty"`
	// 	Currency            string `json:"currency,omitempty"`
	// 	GoodsDetail         []struct {
	// 		GoodsID        string `json:"goods_id,omitempty"`
	// 		Quantity       int64  `json:"quantity,omitempty"`
	// 		UnitPrice      int64  `json:"unit_price,omitempty"`
	// 		DiscountAmount int64  `json:"discount_amount,omitempty"`
	// 		GoodsRemark    string `json:"goods_remark,omitempty"`
	// 	} `json:"goods_detail,omitempty"`
	// } `json:"promotion_detail,omitempty"`
}

/*
	  示例数据

	    {
			"id":"EV-2018022511223320873",
			"create_time":"2018-06-08T10:34:56+08:00",
			"resource_type":"encrypt-resource",
			"event_type":"REFUND.SUCCESS",
			"summary":"退款成功",
			"resource" : {
				"algorithm":"AEAD_AES_256_GCM",
				"original_type": "refund",
				"ciphertext": "...",
				"nonce": "...",
				"associated_data": ""
			}
		}

		{
			"sp_mchid" : "1230000109",
			"sub_mchid" : "1900000109",
			"transaction_id": "1008450740201411110005820873",
			"out_trade_no": "20150806125346",
			"refund_id": "50200207182018070300011301001",
			"out_refund_no": "7752501201407033233368018",
			"refund_status": "SUCCESS",
			"success_time": "2018-06-08T10:34:56+08:00",
			"user_received_account": "招商银行信用卡0403",
			"amount" : {
				"total": 999,
				"refund": 999,
				"payer_total": 999,
				"payer_refund": 999
			}
		}
*/
type WxRefundData struct {
	SpMchid       string `json:"sp_mchid,omitempty"`
	SubMchid      string `json:"sub_mchid,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	OutRefundNo   string `json:"out_refund_no,omitempty"`
	OutTradeNo    string `json:"out_trade_no,omitempty"`
	RefundID      string `json:"refund_id,omitempty"`
	RefundStatus  string `json:"refund_status,omitempty"`
	Amount        struct {
		PayerRefund int64 `json:"payer_refund,omitempty"`
		PayerTotal  int64 `json:"payer_total,omitempty"`
		Refund      int64 `json:"refund,omitempty"`
		Total       int64 `json:"total,omitempty"`
	} `json:"amount,omitempty"`
	SuccessTime         string `json:"success_time,omitempty"`
	UserReceivedAccount string `json:"user_received_account,omitempty"`
}

type ProfitReceiver struct {
	Type         int     `json:"type,omitempty"` // 1-商户 2-个人
	Account      string  `json:"account,omitempty"`
	Name         string  `json:"name,omitempty"`
	SharingRatio float64 `json:"ratio,omitempty"`
}

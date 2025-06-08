package wxutil

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"

	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type PayConf struct {
	MchId                      string
	MchCertificateSerialNumber string
	MchApiv3Key                string
	MchPrivateKeyPath          string
	MchPublicKeyId             string
	MchPublicKeyPath           string
	AppId                      string
	AppSecret                  string
	NotifyUrl                  string
}

type PrepayRequest struct {
	SubMchid        string
	OrderNo         string
	Description     string
	Attach          string // 回调数据
	TotalAmount     int64
	OpenID          string
	IsProfitSharing bool
	TimeExpire      time.Time
}

type PrepayReply struct {
	PrepayID  string `json:"prepay_id"`
	AppID     string `json:"appid"`
	TimeStamp string `json:"time_stamp"`
	NonceStr  string `json:"nonce_str"`
	Package   string `json:"package"`
	SignType  string `json:"sign_type"`
	PaySign   string `json:"pay_sign"`
}

// Prepay 预付请求
func Prepay(ctx context.Context, payConf *PayConf, req *PrepayRequest) (*jsapi.PrepayWithRequestPaymentResponse, error) {

	wechatpayPublicKey, err := utils.LoadPublicKeyWithPath(payConf.MchPublicKeyPath)
	if err != nil {
		return nil, err
	}

	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(payConf.MchPrivateKeyPath)
	if err != nil {
		return nil, err
	}

	opts := []core.ClientOption{
		option.WithWechatPayPublicKeyAuthCipher(
			payConf.MchId,
			payConf.MchCertificateSerialNumber,
			mchPrivateKey,
			payConf.MchPublicKeyId,
			wechatpayPublicKey,
		),
	}

	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("core.NewClient error: %w", err)
	}

	svc := jsapi.JsapiApiService{Client: client}
	resp, _, err := svc.PrepayWithRequestPayment(ctx,
		jsapi.PrepayRequest{
			SpAppid:     core.String(payConf.AppId),
			SpMchid:     core.String(payConf.MchId),
			SubAppid:    nil,
			SubMchid:    core.String(req.SubMchid),
			Description: core.String(req.Description),
			OutTradeNo:  core.String(req.OrderNo),
			Attach:      core.String(req.Attach),
			NotifyUrl:   core.String(payConf.NotifyUrl),
			TimeExpire:  core.Time(req.TimeExpire),
			Amount: &jsapi.Amount{
				Total: core.Int64(req.TotalAmount),
			},
			Payer: &jsapi.Payer{
				SpOpenid: core.String(req.OpenID),
			},
			SettleInfo: &jsapi.SettleInfo{
				ProfitSharing: core.Bool(req.IsProfitSharing),
			},
		},
		payConf.AppId,
	)
	if err != nil {
		return nil, fmt.Errorf("svc.PrepayWithRequestPayment error: %w", err)
	}

	return resp, nil
}

// ParseNotify 解析微信支付回调
func ParseNotify(ctx context.Context, payConf *PayConf, req *http.Request) (map[string]interface{}, error) {

	wechatpayPublicKey, err := utils.LoadPublicKeyWithPath(payConf.MchPublicKeyPath)
	if err != nil {
		return nil, err
	}

	var verifier auth.Verifier
	if payConf.MchId == "1711796550" {
		// 开发环境，使用兼容模式替换平台证书到公钥模式
		mchPrivateKey, err := utils.LoadPrivateKeyWithPath(payConf.MchPrivateKeyPath)
		if err != nil {
			return nil, err
		}
		err = downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mchPrivateKey, payConf.MchCertificateSerialNumber, payConf.MchId, payConf.MchApiv3Key)
		if err != nil {
			//http.Error(w, err.Error(), http.StatusBadRequest)
			return nil, err
		}
		certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(payConf.MchId)
		verifier = verifiers.NewSHA256WithRSACombinedVerifier(certificateVisitor, payConf.MchPublicKeyId, *wechatpayPublicKey)

	} else {
		verifier = verifiers.NewSHA256WithRSAPubkeyVerifier(payConf.MchPublicKeyId, *wechatpayPublicKey)
	}

	handler := notify.NewNotifyHandler(payConf.MchApiv3Key, verifier)

	content := make(map[string]interface{})
	_, err = handler.ParseNotifyRequest(ctx, req, &content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

type QueryPaymentRequest struct {
	SubMchid   string
	OutTradeNo string
}

func QueryPayment(ctx context.Context, payConf *PayConf, req *QueryPaymentRequest) (*partnerpayments.Transaction, error) {
	wechatpayPublicKey, err := utils.LoadPublicKeyWithPath(payConf.MchPublicKeyPath)
	if err != nil {
		return nil, err
	}

	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(payConf.MchPrivateKeyPath)
	if err != nil {
		return nil, err
	}

	opts := []core.ClientOption{
		// option.WithWechatPayAutoAuthCipher(payConf.MchId, payConf.MchCertificateSerialNumber, mchPrivateKey, payConf.MchApiv3Key),
		option.WithWechatPayPublicKeyAuthCipher(
			payConf.MchId,
			payConf.MchCertificateSerialNumber,
			mchPrivateKey,
			payConf.MchPublicKeyId,
			wechatpayPublicKey,
		),
	}

	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	svc := jsapi.JsapiApiService{Client: client}

	trans, _, err := svc.QueryOrderByOutTradeNo(ctx, jsapi.QueryOrderByOutTradeNoRequest{
		SpMchid:    core.String(payConf.MchId),
		SubMchid:   core.String(req.SubMchid),
		OutTradeNo: core.String(req.OutTradeNo),
	})
	if err != nil {
		return nil, fmt.Errorf("QueryOrderByOutTradeNo error: %w", err)
	}

	return trans, nil
}

type RefundRequest struct {
	SubMchId      string
	TransactionId string
	OrderNo       string
	RefundNo      string
	Reason        string
	NotifyUrl     string
	OrderAmount   int64 // 原订单金额数量
	RefundAmount  int64 // 退款金额数量
}

// Refund 退款
func Refund(ctx context.Context, payConf *PayConf, req *RefundRequest) (*refunddomestic.Refund, error) {

	wechatpayPublicKey, err := utils.LoadPublicKeyWithPath(payConf.MchPublicKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(payConf.MchPrivateKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		// option.WithWechatPayAutoAuthCipher(payConf.MchId, payConf.MchCertificateSerialNumber, mchPrivateKey, payConf.MchApiv3Key),
		option.WithWechatPayPublicKeyAuthCipher(
			payConf.MchId,
			payConf.MchCertificateSerialNumber,
			mchPrivateKey,
			payConf.MchPublicKeyId,
			wechatpayPublicKey,
		),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	svc := refunddomestic.RefundsApiService{Client: client}
	resp, _, err := svc.Create(ctx,
		refunddomestic.CreateRequest{
			SubMchid:      core.String(req.SubMchId),
			TransactionId: core.String(req.TransactionId),
			OutTradeNo:    core.String(req.OrderNo),
			OutRefundNo:   core.String(req.RefundNo),
			Reason:        core.String(req.Reason),
			NotifyUrl:     core.String(payConf.NotifyUrl),
			// FundsAccount:  refunddomestic.REQFUNDSACCOUNT_AVAILABLE.Ptr(),
			Amount: &refunddomestic.AmountReq{
				Currency: core.String("CNY"),
				// From: []refunddomestic.FundsFromItem{refunddomestic.FundsFromItem{
				// 	Account: refunddomestic.ACCOUNT_AVAILABLE.Ptr(),
				// 	Amount:  core.Int64(444),
				// }},
				Refund: core.Int64(req.RefundAmount),
				Total:  core.Int64(req.OrderAmount),
			},
			// GoodsDetail: []refunddomestic.GoodsDetail{
			// 	{
			// 		GoodsName:        core.String("iPhone6s 16G"),
			// 		MerchantGoodsId:  core.String("1217752501201407033233368018"),
			// 		RefundAmount:     core.Int64(528800),
			// 		RefundQuantity:   core.Int64(1),
			// 		UnitPrice:        core.Int64(528800),
			// 		WechatpayGoodsId: core.String("1001"),
			// 	},
			// },
		},
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

type QueryRefundRequest struct {
	RefundNo string
	SubMchid string
}

// QueryRefund 查询退款
func QueryRefund(ctx context.Context, payConf *PayConf, req *QueryRefundRequest) (*refunddomestic.Refund, error) {
	wechatpayPublicKey, err := utils.LoadPublicKeyWithPath(payConf.MchPublicKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(payConf.MchPrivateKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		// option.WithWechatPayAutoAuthCipher(payConf.MchId, payConf.MchCertificateSerialNumber, mchPrivateKey, payConf.MchApiv3Key),
		option.WithWechatPayPublicKeyAuthCipher(
			payConf.MchId,
			payConf.MchCertificateSerialNumber,
			mchPrivateKey,
			payConf.MchPublicKeyId,
			wechatpayPublicKey,
		),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	svc := refunddomestic.RefundsApiService{Client: client}
	resp, _, err := svc.QueryByOutRefundNo(ctx,
		refunddomestic.QueryByOutRefundNoRequest{
			OutRefundNo: core.String(req.RefundNo),
			SubMchid:    core.String(req.SubMchid),
		},
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

type OrderReceiverType string

const (
	OrderReceiverType_MerchantID     OrderReceiverType = "MERCHANT_ID"
	OrderReceiverType_PersonalOpenID OrderReceiverType = "PERSONAL_OPENID"
)

type OrderReceiver struct {
	Type        OrderReceiverType // 1、MERCHANT_ID：商户号 2、PERSONAL_OPENID：个人openid（由父商户APPID转换得到）
	Account     string            // 接受者的商户号或openid
	Name        string            // 接受者的商户全称或个人姓名
	Amount      int64             // 分账金额
	Description string            // 描述
}

type ProfitSharingRequest struct {
	SubMchid        string          // 子商户号
	TransactionId   string          // 微信支付单号
	SharingNo       string          // 分账单编号
	UnfreezeUnsplit bool            // 是否解冻剩余未分账余额
	Receivers       []OrderReceiver // 分账接收者
}

// ProfitSharing 请求分账
func ProfitSharing(ctx context.Context, payConf *PayConf, req *ProfitSharingRequest) (*profitsharing.OrdersEntity, error) {
	wechatpayPublicKey, err := utils.LoadPublicKeyWithPath(payConf.MchPublicKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(payConf.MchPrivateKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		// option.WithWechatPayAutoAuthCipher(payConf.MchId, payConf.MchCertificateSerialNumber, mchPrivateKey, payConf.MchApiv3Key),
		option.WithWechatPayPublicKeyAuthCipher(
			payConf.MchId,
			payConf.MchCertificateSerialNumber,
			mchPrivateKey,
			payConf.MchPublicKeyId,
			wechatpayPublicKey,
		),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	var receivers []profitsharing.CreateOrderReceiver
	for _, receiver := range req.Receivers {
		receivers = append(receivers, profitsharing.CreateOrderReceiver{
			Type:        core.String(string(receiver.Type)),
			Account:     core.String(receiver.Account),
			Name:        core.String(receiver.Name),
			Amount:      core.Int64(receiver.Amount),
			Description: core.String(receiver.Description),
		})
	}
	if len(receivers) == 0 {
		return nil, fmt.Errorf("receivers is empty")
	}

	svc := profitsharing.OrdersApiService{Client: client}
	resp, _, err := svc.CreateOrder(ctx,
		profitsharing.CreateOrderRequest{
			Appid:         core.String(payConf.AppId),
			TransactionId: core.String(req.TransactionId),
			OutOrderNo:    core.String(req.SharingNo),
			Receivers:     receivers,
			// SubAppid:        core.String("wx8888888888888889"),
			SubMchid:        core.String(req.SubMchid),
			UnfreezeUnsplit: core.Bool(req.UnfreezeUnsplit),
		},
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

type QueryProfitSharingRequest struct {
	TransactionId string // 微信支付单号
	SubMchid      string // 子商户号
	SharingNo     string // 分账单编号
}

func QueryProfitSharing(ctx context.Context, payConf *PayConf, req *QueryProfitSharingRequest) (*profitsharing.OrdersEntity, error) {
	wechatpayPublicKey, err := utils.LoadPublicKeyWithPath(payConf.MchPublicKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(payConf.MchPrivateKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		// option.WithWechatPayAutoAuthCipher(payConf.MchId, payConf.MchCertificateSerialNumber, mchPrivateKey, payConf.MchApiv3Key),
		option.WithWechatPayPublicKeyAuthCipher(
			payConf.MchId,
			payConf.MchCertificateSerialNumber,
			mchPrivateKey,
			payConf.MchPublicKeyId,
			wechatpayPublicKey,
		),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	svc := profitsharing.OrdersApiService{Client: client}
	resp, _, err := svc.QueryOrder(ctx,
		profitsharing.QueryOrderRequest{
			TransactionId: core.String(req.TransactionId),
			OutOrderNo:    core.String(req.SharingNo),
			SubMchid:      core.String(req.SubMchid),
		},
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

type QueryProfitSharingAmountRequest struct {
	TransactionId string // 微信支付单号
	SharingNo     string // 分账单编号
}

// QueryProfitSharingAmount 查询订单剩余分账数量
func QueryProfitSharingAmount(ctx context.Context, payConf *PayConf, req *QueryProfitSharingAmountRequest) (*profitsharing.QueryOrderAmountResponse, error) {
	wechatpayPublicKey, err := utils.LoadPublicKeyWithPath(payConf.MchPublicKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(payConf.MchPrivateKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		// option.WithWechatPayAutoAuthCipher(payConf.MchId, payConf.MchCertificateSerialNumber, mchPrivateKey, payConf.MchApiv3Key),
		option.WithWechatPayPublicKeyAuthCipher(
			payConf.MchId,
			payConf.MchCertificateSerialNumber,
			mchPrivateKey,
			payConf.MchPublicKeyId,
			wechatpayPublicKey,
		),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	svc := profitsharing.TransactionsApiService{Client: client}
	resp, _, err := svc.QueryOrderAmount(ctx,
		profitsharing.QueryOrderAmountRequest{
			TransactionId: core.String(req.TransactionId),
		},
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

type UnfreezeProfitSharingRequest struct {
	SubMchid      string // 子商户号
	TransactionId string // 微信支付单号
	SharingNo     string // 分账单编号
	Description   string // 描述内容（如：解冻全部剩余资金）
}

// UnfreezeProfitSharing 解冻全部剩余资金
func UnfreezeProfitSharing(ctx context.Context, payConf *PayConf, req *UnfreezeProfitSharingRequest) (*profitsharing.OrdersEntity, error) {
	wechatpayPublicKey, err := utils.LoadPublicKeyWithPath(payConf.MchPublicKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(payConf.MchPrivateKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		// option.WithWechatPayAutoAuthCipher(payConf.MchId, payConf.MchCertificateSerialNumber, mchPrivateKey, payConf.MchApiv3Key),
		option.WithWechatPayPublicKeyAuthCipher(
			payConf.MchId,
			payConf.MchCertificateSerialNumber,
			mchPrivateKey,
			payConf.MchPublicKeyId,
			wechatpayPublicKey,
		),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	svc := profitsharing.OrdersApiService{Client: client}
	resp, _, err := svc.UnfreezeOrder(ctx,
		profitsharing.UnfreezeOrderRequest{
			TransactionId: core.String(req.TransactionId),
			OutOrderNo:    core.String(req.SharingNo),
			SubMchid:      core.String(req.SubMchid),
			Description:   core.String(req.Description),
		},
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

type AddReceiverRequest struct {
	Type         profitsharing.ReceiverType         // 接受者类型：商户ID或个人openid
	Account      string                             // 商户ID或openid
	Name         string                             // 商户全称或个人姓名
	RelationType profitsharing.ReceiverRelationType // 接受者关联类型
	SubMchId     string                             // 子商户号
}

// AddReceiver 添加分账接收方
func AddReceiver(ctx context.Context, payConf *PayConf, req *AddReceiverRequest) (*profitsharing.AddReceiverResponse, error) {
	wechatpayPublicKey, err := utils.LoadPublicKeyWithPath(payConf.MchPublicKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(payConf.MchPrivateKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		// option.WithWechatPayAutoAuthCipher(payConf.MchId, payConf.MchCertificateSerialNumber, mchPrivateKey, payConf.MchApiv3Key),
		option.WithWechatPayPublicKeyAuthCipher(
			payConf.MchId,
			payConf.MchCertificateSerialNumber,
			mchPrivateKey,
			payConf.MchPublicKeyId,
			wechatpayPublicKey,
		),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	svc := profitsharing.ReceiversApiService{Client: client}
	resp, _, err := svc.AddReceiver(ctx,
		profitsharing.AddReceiverRequest{
			Appid:        core.String(payConf.AppId),
			SubMchid:     core.String(req.SubMchId),
			Type:         req.Type.Ptr(),
			Account:      core.String(req.Account),
			Name:         core.String(req.Name),
			RelationType: req.RelationType.Ptr(),
			// CustomRelation: core.String("代理商"),
			// RelationType:   profitsharing.RECEIVERRELATIONTYPE_SERVICE_PROVIDER.Ptr(),
			// SubAppid:       core.String("wx8888888888888889"),
			// SubMchid:       core.String("1900000109"),
		},
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

type DeleteReceiverRequest struct {
	Type     profitsharing.ReceiverType // 接受者类型：商户ID或个人openid
	Account  string                     // 商户ID或openid
	SubMchId string
}

// DeleteReceiver 移除分润接收方
func DeleteReceiver(ctx context.Context, payConf *PayConf, req *DeleteReceiverRequest) (*profitsharing.DeleteReceiverResponse, error) {
	wechatpayPublicKey, err := utils.LoadPublicKeyWithPath(payConf.MchPublicKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(payConf.MchPrivateKeyPath)
	if err != nil {
		return nil, err
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		// option.WithWechatPayAutoAuthCipher(payConf.MchId, payConf.MchCertificateSerialNumber, mchPrivateKey, payConf.MchApiv3Key),
		option.WithWechatPayPublicKeyAuthCipher(
			payConf.MchId,
			payConf.MchCertificateSerialNumber,
			mchPrivateKey,
			payConf.MchPublicKeyId,
			wechatpayPublicKey,
		),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	svc := profitsharing.ReceiversApiService{Client: client}
	resp, _, err := svc.DeleteReceiver(ctx,
		profitsharing.DeleteReceiverRequest{
			Appid:    core.String(payConf.AppId),
			SubMchid: core.String(req.SubMchId),
			Type:     req.Type.Ptr(),
			Account:  core.String(req.Account),
			// SubAppid: core.String("wx8888888888888889"),
			// SubMchid: core.String("1900000109"),
		},
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

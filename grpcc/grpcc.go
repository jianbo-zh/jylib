package grpcc

import (
	"context"
	"crypto/md5"
	"fmt"
	"sync"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/jianbo-zh/jylib/grpcc/filterc"
	"github.com/jianbo-zh/jylib/helpc"
	"github.com/jianbo-zh/jylib/typec"
	"github.com/jianbo-zh/jylib/util/tlsconf"
	"google.golang.org/protobuf/types/known/durationpb"

	clientv3 "go.etcd.io/etcd/client/v3"
	ggrpc "google.golang.org/grpc"

	carauthV1 "github.com/jianbo-zh/jypb/api/carauth/v1"
	carconfigV1 "github.com/jianbo-zh/jypb/api/carconfig/v1"
	carcontrolV1 "github.com/jianbo-zh/jypb/api/carcontrol/v1"
	cardispatchV1 "github.com/jianbo-zh/jypb/api/cardispatch/v1"
	carflightV1 "github.com/jianbo-zh/jypb/api/carflight/v1"
	carmeasureV1 "github.com/jianbo-zh/jypb/api/carmeasure/v1"
	carorderV1 "github.com/jianbo-zh/jypb/api/carorder/v1"
	carproxyV1 "github.com/jianbo-zh/jypb/api/carproxy/v1"
	filestorageV1 "github.com/jianbo-zh/jypb/api/filestorage/v1"
	messageV1 "github.com/jianbo-zh/jypb/api/message/v1"
	parkmapV1 "github.com/jianbo-zh/jypb/api/parkmap/v1"
)

type IClient interface {
	CarAuthGrpc() ICarAuth
	CarConfigGrpc() ICarConfig
	CarDispatchGrpc() ICarDispatch
	CarOrderGrpc() ICarOrder
	CarProxyGrpc() ICarProxy
	CarControlGrpc() ICarProxy
	CarMeasureGrpc() ICarMeasure
	CarFlightGrpc() ICarFlight
	FileStorageGrpc() IFileStorage
	MessageGrpc() IMessage
	ParkMapGrpc() IParkMap

	CarAuthClient(context.Context, ...filterc.Filter) (carauthV1.CarAuthClient, error)
	CarConfigClient(context.Context, ...filterc.Filter) (carconfigV1.ConfigClient, error)
	CarDispatchClient(context.Context, ...filterc.Filter) (cardispatchV1.DispatchClient, error)
	CarOrderClient(context.Context, ...filterc.Filter) (carorderV1.CarOrderClient, error)
	CarProxyClient(context.Context, ...filterc.Filter) (carproxyV1.CarProxyClient, error)
	CarControlClient(context.Context, ...filterc.Filter) (carcontrolV1.CarControlClient, error)
	CarMeasureClient(context.Context, ...filterc.Filter) (carmeasureV1.CarMeasureClient, error)
	CarFlightClient(context.Context, ...filterc.Filter) (carflightV1.CarFlightClient, error)
	FileStorageClient(context.Context, ...filterc.Filter) (filestorageV1.FileStorageClient, error)
	MessageClient(context.Context, ...filterc.Filter) (messageV1.MessageClient, error)
	ParkMapClient(context.Context, ...filterc.Filter) (parkmapV1.ParkMapClient, error)
}

type Client struct {
	env       string
	discovery *etcd.Registry
	logger    log.Logger
	CaCrtFile string

	grpcConns sync.Map
}

type TlsConf struct {
	CaCrt  string
	SvrKey string
	SvrCrt string
}

type EtcdConf struct {
	Username    string
	Password    string
	Endpoints   []string
	TlsConf     *TlsConf
	DialTimeout *durationpb.Duration
}

func NewClient(env string, caCrtFile string, etcdConf *EtcdConf, logger log.Logger) IClient {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdConf.Endpoints,
		Username:    etcdConf.Username,
		Password:    etcdConf.Password,
		TLS:         tlsconf.NewServerTlsConfig(etcdConf.TlsConf.SvrKey, etcdConf.TlsConf.SvrCrt, etcdConf.TlsConf.CaCrt),
		DialTimeout: etcdConf.DialTimeout.AsDuration(),
	})
	if err != nil {
		panic(err)
	}
	discovery := etcd.New(client)

	// 全局 balancer name 的方式来注入 selector
	selector.SetGlobalSelector(wrr.NewBuilder())

	return &Client{
		env:       env,
		discovery: discovery,
		logger:    logger,
		CaCrtFile: caCrtFile,
	}
}

// CarAuthGrpc
func (c *Client) CarAuthGrpc() ICarAuth { return NewCarAuthGrpc(c) }

// CarConfigGrpc
func (c *Client) CarConfigGrpc() ICarConfig { return NewCarConfigGrpc(c) }

// CarDispatchGrpc
func (c *Client) CarDispatchGrpc() ICarDispatch { return NewCarDispatchGrpc(c) }

// CarOrderGrpc
func (c *Client) CarOrderGrpc() ICarOrder { return NewCarOrderGrpc(c) }

// CarProxyGrpc
func (c *Client) CarProxyGrpc() ICarProxy { return NewCarProxyGrpc(c) }

// CarProxyGrpc
func (c *Client) CarControlGrpc() ICarProxy { return NewCarControlGrpc(c) }

// CarMeasureGrpc
func (c *Client) CarMeasureGrpc() ICarMeasure { return NewCarMeasureGrpc(c) }

// CarFlightGrpc
func (c *Client) CarFlightGrpc() ICarFlight { return NewCarFlightGrpc(c) }

// FileStorageGrpc
func (c *Client) FileStorageGrpc() IFileStorage { return NewFileStorageGrpc(c) }

// MessageGrpc
func (c *Client) MessageGrpc() IMessage { return NewMessageGrpc(c) }

// ParkMapGrpc
func (c *Client) ParkMapGrpc() IParkMap { return NewParkMapGrpc(c) }

// CarAuthClient
func (c *Client) CarAuthClient(ctx context.Context, filters ...filterc.Filter) (carauthV1.CarAuthClient, error) {
	conn, err := c.grpcConn(ctx, typec.Service_MsCarAuth, filters...)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return carauthV1.NewCarAuthClient(conn), nil
}

// CarConfigClient
func (c *Client) CarConfigClient(ctx context.Context, filters ...filterc.Filter) (carconfigV1.ConfigClient, error) {
	conn, err := c.grpcConn(ctx, typec.Service_MsCarConfig, filters...)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return carconfigV1.NewConfigClient(conn), nil
}

// CarDispatchClient
func (c *Client) CarDispatchClient(ctx context.Context, filters ...filterc.Filter) (cardispatchV1.DispatchClient, error) {
	conn, err := c.grpcConn(ctx, typec.Service_MsCarDispatch, filters...)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return cardispatchV1.NewDispatchClient(conn), nil
}

// CarOrderClient
func (c *Client) CarOrderClient(ctx context.Context, filters ...filterc.Filter) (carorderV1.CarOrderClient, error) {
	conn, err := c.grpcConn(ctx, typec.Service_MsOrder, filters...)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return carorderV1.NewCarOrderClient(conn), nil
}

// CarProxyClient
func (c *Client) CarProxyClient(ctx context.Context, filters ...filterc.Filter) (carproxyV1.CarProxyClient, error) {
	conn, err := c.grpcConn(ctx, typec.Service_GwCarProxy, filters...)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return carproxyV1.NewCarProxyClient(conn), nil
}

// CarControlClient
func (c *Client) CarControlClient(ctx context.Context, filters ...filterc.Filter) (carcontrolV1.CarControlClient, error) {
	conn, err := c.grpcConn(ctx, typec.Service_MsCarControl, filters...)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return carcontrolV1.NewCarControlClient(conn), nil
}

// CarProxyClient
func (c *Client) CarMeasureClient(ctx context.Context, filters ...filterc.Filter) (carmeasureV1.CarMeasureClient, error) {
	conn, err := c.grpcConn(ctx, typec.Service_MsCarMeasure, filters...)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return carmeasureV1.NewCarMeasureClient(conn), nil
}

// CarFlightClient
func (c *Client) CarFlightClient(ctx context.Context, filters ...filterc.Filter) (carflightV1.CarFlightClient, error) {
	conn, err := c.grpcConn(ctx, typec.Service_MsCarFlight, filters...)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return carflightV1.NewCarFlightClient(conn), nil
}

// FileStorageClient
func (c *Client) FileStorageClient(ctx context.Context, filters ...filterc.Filter) (filestorageV1.FileStorageClient, error) {
	conn, err := c.grpcConn(ctx, typec.Service_MsFileStorage, filters...)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return filestorageV1.NewFileStorageClient(conn), nil
}

// MessageClient
func (c *Client) MessageClient(ctx context.Context, filters ...filterc.Filter) (messageV1.MessageClient, error) {
	conn, err := c.grpcConn(ctx, typec.Service_MsMessage, filters...)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return messageV1.NewMessageClient(conn), nil
}

// ParkMapClient
func (c *Client) ParkMapClient(ctx context.Context, filters ...filterc.Filter) (parkmapV1.ParkMapClient, error) {
	conn, err := c.grpcConn(ctx, typec.Service_MsParkMap, filters...)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return parkmapV1.NewParkMapClient(conn), nil
}

func (c *Client) grpcConn(ctx context.Context, service string, filters ...filterc.Filter) (*ggrpc.ClientConn, error) {
	sumRaw := service
	for _, filter := range filters {
		sumRaw += filter.Uuid()
	}
	uuid := fmt.Sprintf("%x", md5.Sum([]byte(sumRaw)))
	value, ok := c.grpcConns.Load(uuid)
	if !ok {
		conn, err := c.grpcDial(ctx, service, filters...)
		if err != nil {
			return nil, fmt.Errorf("grpc.Dial error: %w", err)
		}

		c.grpcConns.Store(uuid, conn)
		return conn, nil
	}
	conn, ok := value.(*ggrpc.ClientConn)
	if !ok {
		return nil, fmt.Errorf("value assert *grpc.ClientConn failed")
	}
	return conn, nil
}

func (c *Client) grpcDial(ctx context.Context, serverName string, filters ...filterc.Filter) (*ggrpc.ClientConn, error) {
	options := []grpc.ClientOption{
		grpc.WithDiscovery(c.discovery),
		grpc.WithEndpoint(helpc.ServerEndpoint(c.env, serverName)),
		grpc.WithTimeout(5 * time.Second),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
			metadata.Client(),
		),
		grpc.WithLogger(c.logger),
		// 使用 DialInsecure 再加上 WithTLSConfig 会报错
		grpc.WithTLSConfig(
			tlsconf.NewClientTlsConfig("", "", c.CaCrtFile),
		),
	}
	if filters != nil {
		for _, filter := range filters {
			options = append(options, grpc.WithNodeFilter(filter.Filter()))
		}
	}
	return grpc.Dial(ctx, options...)
}

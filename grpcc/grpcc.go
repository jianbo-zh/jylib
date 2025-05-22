package grpcc

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/jianbo-zh/jylib/helpc"
	"github.com/jianbo-zh/jylib/typec"
	"github.com/jianbo-zh/jylib/util/tlsconf"
	"google.golang.org/protobuf/types/known/durationpb"

	clientv3 "go.etcd.io/etcd/client/v3"

	authV1 "github.com/jianbo-zh/jypb/api/carauth/v1"
	configV1 "github.com/jianbo-zh/jypb/api/carconfig/v1"
	dispatchV1 "github.com/jianbo-zh/jypb/api/cardispatch/v1"
	orderV1 "github.com/jianbo-zh/jypb/api/carorder/v1"
	proxyV1 "github.com/jianbo-zh/jypb/api/carproxy/v1"
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
	CarMeasureGrpc() ICarMeasure
	FileStorageGrpc() IFileStorage
	MessageGrpc() IMessage
	ParkMapGrpc() IParkMap

	CarAuthClient(context.Context) (authV1.CarAuthClient, error)
	CarConfigClient(context.Context) (configV1.ConfigClient, error)
	CarDispatchClient(context.Context) (dispatchV1.DispatchClient, error)
	CarOrderClient(context.Context) (orderV1.CarOrderClient, error)
	CarProxyClient(context.Context) (proxyV1.CarProxyClient, error)
	CarMeasureClient(context.Context) (proxyV1.CarMeasureClient, error)
	FileStorageClient(context.Context) (filestorageV1.FileStorageClient, error)
	MessageClient(context.Context) (messageV1.MessageClient, error)
	ParkMapClient(context.Context) (parkmapV1.ParkMapClient, error)
}

type Client struct {
	env       string
	discovery *etcd.Registry
	logger    log.Logger
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

func NewClient(env string, etcdConf *EtcdConf, logger log.Logger) IClient {
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

	return &Client{
		env:       env,
		discovery: discovery,
		logger:    logger,
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

// CarMeasureGrpc
func (c *Client) CarMeasureGrpc() ICarMeasure { return NewCarMeasureGrpc(c) }

// FileStorageGrpc
func (c *Client) FileStorageGrpc() IFileStorage { return NewFileStorageGrpc(c) }

// MessageGrpc
func (c *Client) MessageGrpc() IMessage { return NewMessageGrpc(c) }

// ParkMapGrpc
func (c *Client) ParkMapGrpc() IParkMap { return NewParkMapGrpc(c) }

// CarAuthClient
func (c *Client) CarAuthClient(ctx context.Context) (authV1.CarAuthClient, error) {
	conn, err := grpc.Dial(ctx,
		grpc.WithDiscovery(c.discovery),
		grpc.WithEndpoint(helpc.ServerEndpoint(c.env, typec.Service_CarAuth)),
		grpc.WithTimeout(5*time.Second),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
		grpc.WithLogger(c.logger),
		grpc.WithTLSConfig(
			tlsconf.NewClientTlsConfig("", "", ""),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return authV1.NewCarAuthClient(conn), nil
}

// CarConfigClient
func (c *Client) CarConfigClient(ctx context.Context) (configV1.ConfigClient, error) {
	conn, err := grpc.Dial(ctx,
		grpc.WithDiscovery(c.discovery),
		grpc.WithEndpoint(helpc.ServerEndpoint(c.env, typec.Service_CarConfig)),
		grpc.WithTimeout(5*time.Second),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
		grpc.WithLogger(c.logger),
		grpc.WithTLSConfig(
			tlsconf.NewClientTlsConfig("", "", ""),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return configV1.NewConfigClient(conn), nil
}

// CarDispatchClient
func (c *Client) CarDispatchClient(ctx context.Context) (dispatchV1.DispatchClient, error) {
	conn, err := grpc.Dial(ctx,
		grpc.WithDiscovery(c.discovery),
		grpc.WithEndpoint(helpc.ServerEndpoint(c.env, typec.Service_Dispatch)),
		grpc.WithTimeout(5*time.Second),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
		grpc.WithLogger(c.logger),
		grpc.WithTLSConfig(
			tlsconf.NewClientTlsConfig("", "", ""),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return dispatchV1.NewDispatchClient(conn), nil
}

// CarOrderClient
func (c *Client) CarOrderClient(ctx context.Context) (orderV1.CarOrderClient, error) {
	conn, err := grpc.Dial(ctx,
		grpc.WithDiscovery(c.discovery),
		grpc.WithEndpoint(helpc.ServerEndpoint(c.env, typec.Service_CarOrder)),
		grpc.WithTimeout(5*time.Second),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
		grpc.WithLogger(c.logger),
		grpc.WithTLSConfig(
			tlsconf.NewClientTlsConfig("", "", ""),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return orderV1.NewCarOrderClient(conn), nil
}

// CarProxyClient
func (c *Client) CarProxyClient(ctx context.Context) (proxyV1.CarProxyClient, error) {
	conn, err := grpc.Dial(ctx,
		grpc.WithDiscovery(c.discovery),
		grpc.WithEndpoint(helpc.ServerEndpoint(c.env, typec.Service_CarProxy)),
		grpc.WithTimeout(5*time.Second),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
		grpc.WithLogger(c.logger),
		grpc.WithTLSConfig(
			tlsconf.NewClientTlsConfig("", "", ""),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return proxyV1.NewCarProxyClient(conn), nil
}

// CarProxyClient
func (c *Client) CarMeasureClient(ctx context.Context) (proxyV1.CarMeasureClient, error) {
	conn, err := grpc.Dial(ctx,
		grpc.WithDiscovery(c.discovery),
		grpc.WithEndpoint(helpc.ServerEndpoint(c.env, typec.Service_CarProxy)),
		grpc.WithTimeout(5*time.Second),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
		grpc.WithLogger(c.logger),
		grpc.WithTLSConfig(
			tlsconf.NewClientTlsConfig("", "", ""),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return proxyV1.NewCarMeasureClient(conn), nil
}

// FileStorageClient
func (c *Client) FileStorageClient(ctx context.Context) (filestorageV1.FileStorageClient, error) {
	conn, err := grpc.Dial(ctx,
		grpc.WithDiscovery(c.discovery),
		grpc.WithEndpoint(helpc.ServerEndpoint(c.env, typec.Service_FileStorage)),
		grpc.WithTimeout(5*time.Second),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
		grpc.WithLogger(c.logger),
		grpc.WithTLSConfig(
			tlsconf.NewClientTlsConfig("", "", ""),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return filestorageV1.NewFileStorageClient(conn), nil
}

// MessageClient
func (c *Client) MessageClient(ctx context.Context) (messageV1.MessageClient, error) {
	conn, err := grpc.Dial(ctx,
		grpc.WithDiscovery(c.discovery),
		grpc.WithEndpoint(helpc.ServerEndpoint(c.env, typec.Service_Message)),
		grpc.WithTimeout(5*time.Second),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
		grpc.WithLogger(c.logger),
		grpc.WithTLSConfig(
			tlsconf.NewClientTlsConfig("", "", ""),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return messageV1.NewMessageClient(conn), nil
}

// ParkMapClient
func (c *Client) ParkMapClient(ctx context.Context) (parkmapV1.ParkMapClient, error) {
	conn, err := grpc.Dial(ctx,
		grpc.WithDiscovery(c.discovery),
		grpc.WithEndpoint(helpc.ServerEndpoint(c.env, typec.Service_ParkMap)),
		grpc.WithTimeout(5*time.Second),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
		grpc.WithLogger(c.logger),
		grpc.WithTLSConfig(
			tlsconf.NewClientTlsConfig("", "", ""),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial error: %w", err)
	}

	return parkmapV1.NewParkMapClient(conn), nil
}

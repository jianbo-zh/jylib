package tlsconf

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
)

// 用来控制客户端是否证书和服务器主机名。如果设置为true,则不会校验证书以及证书中的主机名和服务器主机名是否一致。
const insecureSkipVerify = true

// TLSInfo holds the SSL certificates paths.
type TLSInfo struct {
	CertFile           string `json:"CertFile"`
	KeyFile            string `json:"KeyFile"`
	CAFile             string `json:"CAFile"`
	InsecureSkipVerify bool   `json:"InsecureSkipVerify"`
}

func (info TLSInfo) Scheme() string {
	if info.KeyFile != "" && info.CertFile != "" {
		return "https"
	} else {
		return "http"
	}
}

func (info TLSInfo) ServerConfig() (*tls.Config, error) {
	if info.KeyFile == "" || info.CertFile == "" {
		return nil, fmt.Errorf("KeyFile and CertFile must both be present[key: %v, cert: %v]", info.KeyFile, info.CertFile)
	}

	var cfg tls.Config
	cfg.InsecureSkipVerify = info.InsecureSkipVerify
	//cfg.ServerName = "host.docker.internal"
	cfg.MinVersion = tls.VersionSSL30

	tlsCert, err := tls.LoadX509KeyPair(info.CertFile, info.KeyFile)
	if err != nil {
		log.Println("LoadX509KeyPair error:", err)
		return nil, err
	}

	cfg.Certificates = []tls.Certificate{tlsCert}

	if info.CAFile != "" {
		cfg.ClientAuth = tls.RequireAndVerifyClientCert
		cp, err := newCertPool(info.CAFile)
		if err != nil {
			log.Fatalln("read cert file error:", err)
			return nil, err
		}

		cfg.RootCAs = cp
		cfg.ClientCAs = cp
	} else {
		cfg.ClientAuth = tls.NoClientCert
	}

	return &cfg, nil
}

func (info TLSInfo) ClientConfig() (*tls.Config, error) {
	var cfg tls.Config
	cfg.InsecureSkipVerify = info.InsecureSkipVerify
	//cfg.ServerName = "host.docker.internal"
	cfg.MinVersion = tls.VersionSSL30

	if info.CAFile != "" {
		cp, err := newCertPool(info.CAFile)
		if err != nil {
			log.Fatalln("read cert file error:", err)
			return nil, err
		}

		cfg.RootCAs = cp
	}

	if info.KeyFile == "" || info.CertFile == "" {
		return &cfg, nil
	}

	tlsCert, err := tls.LoadX509KeyPair(info.CertFile, info.KeyFile)
	if err != nil {
		log.Fatalln("read pair file error:", err)
		return nil, err
	}

	cfg.Certificates = []tls.Certificate{tlsCert}

	return &cfg, nil
}

// newCertPool creates x509 certPool with provided CA file
func newCertPool(caFile string) (*x509.CertPool, error) {
	certPool := x509.NewCertPool()
	pemByte, err := os.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	if !certPool.AppendCertsFromPEM(pemByte) {
		return nil, fmt.Errorf("can't add CA cert")
	}
	return certPool, nil
}

// NewServerTlsConfig 创建服务端TLS证书认证配置
// param [in] keyFile 服务端私钥文件路径，必须提供
// param [in] certFile 服务端证书文件路径，必须提供
// param [in] caFile CA根证书，如果提供则为双向认证，否则为单向认证
func NewServerTlsConfig(keyFile, certFile, caFile string) *tls.Config {
	var info = TLSInfo{
		CertFile:           certFile,
		KeyFile:            keyFile,
		CAFile:             caFile,
		InsecureSkipVerify: insecureSkipVerify,
	}
	cfg, _ := info.ServerConfig()
	return cfg
}

// NewClientTlsConfig 创建客户端端TLS证书认证配置
// param [in] keyFile 客户端私钥文件路径
// param [in] certFile 客户端证书文件路径
// param [in] caFile CA根证书
func NewClientTlsConfig(keyFile, certFile, caFile string) *tls.Config {
	var info = TLSInfo{
		CertFile:           certFile,
		KeyFile:            keyFile,
		CAFile:             caFile,
		InsecureSkipVerify: insecureSkipVerify,
	}
	cfg, _ := info.ClientConfig()
	return cfg
}

package crypto

import "github.com/wenzhenxi/gorsa"

// PublicEncrypt 公钥加密
func PublicEncrypt(data string, publicKey string) (string, error) {
	return gorsa.PublicEncrypt(data, publicKey)
}

// PublicDecrypt 公钥解密
func PublicDecrypt(data string, publicKey string) (string, error) {
	return gorsa.PublicDecrypt(data, publicKey)
}

// PrivateEncrypt 私钥加密
func PrivateEncrypt(data string, privateKey string) (string, error) {
	return gorsa.PriKeyEncrypt(data, privateKey)
}

// PrivateDecrypt 私钥解密
func PrivateDecrypt(data string, privateKey string) (string, error) {
	return gorsa.PriKeyDecrypt(data, privateKey)
}

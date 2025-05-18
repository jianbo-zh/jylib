package crypto

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func DesCBCEncrypt(origData []byte, key [8]byte) ([]byte, error) {
	key1 := key[:]
	block, err := des.NewCipher(key1)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key1)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func DesCBCDecrypt(crypted []byte, key [8]byte) ([]byte, error) {
	key1 := key[:]
	block, err := des.NewCipher(key1)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key1)
	//origData := make([]byte, len(crypted))
	origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	//origData = PKCS5UnPadding(origData)

	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

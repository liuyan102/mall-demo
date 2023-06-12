package util

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"errors"
)

// AesEncrypt 对称加密
func AesEncrypt(money string, key string) (string, error) {
	// 创建加密实例
	cipher, err := aes.NewCipher([]byte(key)) // key必须为16
	if err != nil {
		return "", err
	}
	// 判断加密块的大小
	size := cipher.BlockSize()
	// pkcs7填充
	encryptBytes := pkcs7Padding([]byte(money), size)
	// 初始化加密数据接收切片
	out := make([]byte, len(encryptBytes))
	// 执行加密
	cipher.Encrypt(out, encryptBytes)
	// 将字节编码为字符串
	return hex.EncodeToString(out), nil
}

// pkcs7填充
func pkcs7Padding(data []byte, size int) []byte {
	// 判断缺少几位长度 最少1位 最多size
	padding := size - len(data)%size
	// 补足位数 把切片[]byte{byte(padding)}再复制padding个
	padByte := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padByte...)
}

// AesDecrypt 解密
func AesDecrypt(money string, key string) (string, error) {
	// 解码
	decryptBytes, err := hex.DecodeString(money)
	if err != nil {
		return "", err
	}
	// 创建实例
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	// 初始化解密数据接收切片
	out := make([]byte, len(decryptBytes))
	// 解密
	cipher.Decrypt(out, decryptBytes)
	// 去除填充
	out, err = pkcs7UnPadding(out)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误")
	}
	// 获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

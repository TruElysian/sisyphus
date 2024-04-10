package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// MARK: AES-GCM 加密

// 调用方式：使用 generate.GenerateSafeRandomBytes 生成对应密钥再调用加解密函数，可以参照测试代码

// 使用 AES-GCM 算法对明文进行加密 传入 plainText string and key base64 string
func AESGCMEncryptData(plainText, key string) (string, error) {
	// 将密钥和明文转换为字节数组
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	plainTextBytes := []byte(plainText)
	cipherText, err := AESGCMEncrypt(plainTextBytes, keyBytes)

	return base64.StdEncoding.EncodeToString(cipherText), err
}

// 使用 AES-GCM 算法对明文进行加密
func AESGCMEncrypt(plainText, key []byte) ([]byte, error) {
	// 创建一个新的 AES 密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建一个 GCM 模式的块
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 生成随机的 Nonce 值，必须是 12 字节长
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// 对明文进行加密
	cipherText := gcm.Seal(nil, nonce, plainText, nil)
	// 将 Nonce 值附加到加密后的数据中
	cipherText = append(nonce, cipherText...)
	return cipherText, nil
}

// MARK: AES-GCM 解密

// 使用 AES-GCM 算法对密文进行解密 传入 cipherText base64 string and key base64 string
func AESGCMDecryptData(cipherText, key string) (string, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	plainText, err := AESGCMDecrypt(cipherTextBytes, keyBytes)
	return string(plainText), err
}

// 使用 AES-GCM 算法对密文进行解密
func AESGCMDecrypt(cipherText, key []byte) ([]byte, error) {
	// 创建一个新的 AES 密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建一个 GCM 模式的块
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 检查密文长度是否足够
	if len(cipherText) < gcm.NonceSize() {
		return nil, errors.New("cipherText too short")
	}

	// 从密文中提取 Nonce 值
	nonce := cipherText[:gcm.NonceSize()]
	// 使用 Nonce 值和密文进行解密
	plainText, err := gcm.Open(nil, nonce, cipherText[gcm.NonceSize():], nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

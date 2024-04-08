package generate

import "crypto/rand"

// 健壮的生成随机数字节数组
func GenerateSafeRandomBytes(length uint32) ([]byte, error) {
	random := make([]byte, length)
	_, err := rand.Read(random) // 使用 crypto/rand 来生成
	if err != nil {
		return nil, err
	}
	return random, nil
}

package util

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// 生成一个带时间戳的 token
func GenerateTokenWithTimestamp(username string) string {
	// 生成当前时间戳
	timestamp := time.Now().UnixNano()

	// 对用户名和时间戳进行 SHA-256 哈希处理
	hasher := sha256.New()
	hasher.Write([]byte(username + strconv.FormatInt(timestamp, 10)))
	hashBytes := hasher.Sum(nil)

	// 将哈希值转换为十六进制字符串
	hashString := hex.EncodeToString(hashBytes)

	// 截取前 16 个字符作为 token
	token := hashString[:16]

	return token
}

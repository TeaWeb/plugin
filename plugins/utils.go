package plugins

import (
	"math/rand"
	"time"
)

// 取得随机字符串
// 代码来自 https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func RandString(n int) string {
	const randomLetterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	const (
		randomLetterIdxBits = 6                          // 6 bits to represent a letter index
		randomLetterIdxMask = 1<<randomLetterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		randomLetterIdxMax  = 63 / randomLetterIdxBits   // # of letter indices fitting in 63 bits
	)

	var src = rand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), randomLetterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), randomLetterIdxMax
		}
		if idx := int(cache & randomLetterIdxMask); idx < len(randomLetterBytes) {
			b[i] = randomLetterBytes[idx]
			i--
		}
		cache >>= randomLetterIdxBits
		remain--
	}

	return string(b)
}

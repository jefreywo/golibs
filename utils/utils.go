package utils

import (
	"math"
	"math/rand"
	"regexp"
	"time"
)

// emoji Unicode参考网址
// https://www.unicode.org/Public/emoji/13.0/emoji-sequences.txt
// https://www.unicode.org/Public/emoji/13.0/emoji-zwj-sequences.txt
// https://apps.timwhitlock.info/emoji/tables/unicode#block-1-emoticons

// 通过正则匹配过滤字符串中的emoji颜文字
func FilterEmoji(str string) string {
	var emojiRx = regexp.MustCompile(`[\x{1F000}-\x{1F6FF}|[\x{2600}-\x{26FF}]`)
	return emojiRx.ReplaceAllString(str, ``)
}

// 四舍五入保留n位小数
func Round(x float64, n int) float64 {
	return math.Round(x*math.Pow10(n)) / math.Pow10(n)
}

// 两个floa64的数是否在误差范围内相等，
func IsEqual(x, y float64, precision float64) bool {
	return math.Dim(x, y) < precision
}

// 随机返回数组中的某值
func RandArray(arr ...interface{}) interface{} {
	rand.Seed(time.Now().UnixNano())
	return arr[rand.Intn(len(arr))]
}

// 获取[a,b)范围内的随机整数，要求 a < b
func RandInterval(a, b int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(b-a+1) + a
}

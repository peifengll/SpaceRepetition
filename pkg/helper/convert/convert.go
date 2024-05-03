package convert

import (
	"log"
	"regexp"
	"strconv"
)

const (
	base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func IntToBase62(n int) string {
	if n == 0 {
		return string(base62[0])
	}

	var result []byte
	for n > 0 {
		result = append(result, base62[n%62])
		n /= 62
	}

	// 反转字符串
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func ParseStringToFloatSlice(s string) ([]float64, error) {
	// 使用正则表达式匹配数字和小数点
	re := regexp.MustCompile(`\d+\.\d+|\d+`)
	matches := re.FindAllString(s, -1)

	// 将匹配到的字符串转换为浮点数并添加到切片中
	floatSlice := make([]float64, len(matches))
	for i, match := range matches {
		f, err := strconv.ParseFloat(match, 64)
		if err != nil {
			log.Println("Error:", err)
			return nil, err
		}
		floatSlice[i] = f
	}
	return floatSlice, nil
}

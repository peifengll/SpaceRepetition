package server

import (
	"fmt"
	"testing"
	"time"
)

func TestDateSub(t *testing.T) {
	// 定义两个日期
	date1 := time.Date(2024, 5, 3, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2024, 5, 2, 23, 0, 0, 0, time.UTC)
	// 将时分秒截断
	date1 = date1.Truncate(24 * time.Hour)
	date2 = date2.Truncate(24 * time.Hour)

	// 计算日期差异
	difference := date1.Sub(date2)
	// 将时间差异转换为天数
	daysDifference := int(difference.Hours() / 24)

	// 输出结果
	fmt.Printf("日期1: %s\n", date1.Format("2006-01-02"))
	fmt.Printf("日期2: %s\n", date2.Format("2006-01-02"))
	fmt.Printf("日期差异: %d 天\n", daysDifference)
}

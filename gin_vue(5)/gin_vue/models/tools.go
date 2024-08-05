package models

import "time"

// GetUnix 获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// 获取当前的目期
func GetDate() string {
	template := "2006-01-02"
	return time.Now().Format(template)
}

// 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// 时间戳转换成日期
func UnixToTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

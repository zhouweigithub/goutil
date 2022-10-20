package timeutil

import (
	"fmt"
	"time"
)

// 日期格式化字符串
var DateFormatString = "2006-01-02"

// 时间格式化字符串
var TimeFormatString = "15:04:05"

// 时间格式化字符串
var DateTimeFormatString = "2006-01-02 15:04:05"

// 取当前日期
func GetCurrentDate() time.Time {
	year, month, day := time.Now().Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// 取当前日期字符串
func GetCurrentDateString() string {
	return GetCurrentDate().Format(DateFormatString)
}

// 格式化日期
func FormateDateString(date time.Time) string {
	return date.Format(DateFormatString)
}

// 格式化时间
func FormatTimeString(time time.Time) string {
	return time.Format(TimeFormatString)
}

// 格式化日期与时间
func FormateDateTimeString(date time.Time) string {
	return date.Format(DateTimeFormatString)
}

// 字符串转本地时间
func ConvertStringToLocalTime(layout string, timeString string) (time.Time, error) {
	return time.ParseInLocation(layout, timeString, time.Local)
}

// 格式化时间差为时分秒
func FormatTimeDuringString(totalSeconds float64) string {
	totalInt := int(totalSeconds)
	hour := totalInt / 60 / 60
	minute := (totalInt - hour*60*60) / 60
	second := totalInt % 60
	mic := totalSeconds - float64(totalInt)
	return fmt.Sprintf("%d 小时 %d 分 %.2f 秒", hour, minute, float64(second)+mic)
}

// 日期转换为时间戳
func ConvertToTimeStamp(datetime time.Time) int64 {
	return datetime.Unix()
}

// 时间戳转换为日期
func ConvertToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

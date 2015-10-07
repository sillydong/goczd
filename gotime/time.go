package gotime

import (
	"fmt"
	"time"
)

const (
	FORMAT_YYYY_MM_DD          = "2006-01-02"
	FORMAT_YYYY_MM_DD_HH_II_SS = "2006-01-02 15:04:05"
	Y                          = "2006"
	M                          = "01"
	D                          = "02"
	H                          = "15"
	I                          = "04"
	S                          = "05"
)

//获取当前时间戳
func GetTimestamp() int64 {
	return time.Now().Unix()
}

//按照指定格式显示当前时间
func GetTimeStr(format string) string {
	return time.Now().Format(format)
}

//字符串转时间戳
func StrToTime(str, format string) (int64, error) {
	timetime, err := time.Parse(format, str)
	if err != nil {
		return 0, err
	} else {
		return timetime.Unix(), nil
	}
}

//时间戳转字符串
func TimeToStr(timestamp int64, format string) string {
	return time.Unix(timestamp, 0).Format(format)
}

//友好的时间显示
func FriendlyTime(timestamp int64) string {
	current := time.Now()
	diff := current.Unix() - timestamp

	switch {
	case diff > 0 && diff < 60:
		return fmt.Sprintf("%d秒前", diff)
	case diff >= 60 && diff < 3600:
		return fmt.Sprintf("%d分钟前", diff/60)
	case diff >= 3600 && diff < 86400:
		return fmt.Sprintf("%d小时前", diff/3600)
	case diff >= 86400 && diff < 86400*2:
		return "昨天"
	case diff >= 86400*2 && diff < 86400*3:
		return "前天"
	default:
		return TimeToStr(timestamp, Y+"-"+M+"-"+D)
	}
}

package gotime

import (
	"fmt"
	"strconv"
	"time"
)

const (
	FORMAT_YYYY_MM             = "2006-01"
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

//显示指定月份前后一个月份
func CalcMonth(yearmonth string) (string, string, string, error) {
	if len(yearmonth) != 6 {
		return "", "", "", fmt.Errorf("%v", "yearmonth format should be 201601")
	}
	year := yearmonth[0:4]
	month := yearmonth[4:6]
	iyear, _ := strconv.Atoi(year)
	switch month {
	case "01":
		return strconv.Itoa(iyear-1) + "12", year + "01", year + "02", nil
	case "02":
		return year + "01", year + "02", year + "03", nil
	case "03":
		return year + "02", year + "03", year + "04", nil
	case "04":
		return year + "03", year + "04", year + "05", nil
	case "05":
		return year + "04", year + "05", year + "06", nil
	case "06":
		return year + "05", year + "06", year + "07", nil
	case "07":
		return year + "06", year + "07", year + "08", nil
	case "08":
		return year + "07", year + "08", year + "09", nil
	case "09":
		return year + "08", year + "09", year + "10", nil
	case "10":
		return year + "09", year + "10", year + "11", nil
	case "11":
		return year + "10", year + "11", year + "12", nil
	case "12":
		return year + "11", year + "12", strconv.Itoa(iyear+1) + "01", nil
	}
	return "", "", "", fmt.Errorf("%v", "yearmonth format should be 201601")
}

package gotime

import "testing"

func Test_GetTimestamp(t *testing.T) {
	t.Log(GetTimestamp())
}

func Test_StrToTime(t *testing.T) {
	timestamp := GetTimestamp()
	t.Log(TimeToStr(timestamp, FORMAT_YYYY_MM_DD))
	t.Log(TimeToStr(timestamp, FORMAT_YYYY_MM_DD_HH_II_SS))
}

func Test_TimeToStr(t *testing.T) {
	timestr := "2015-09-01 00:00:00"
	timestamp, err := StrToTime(timestr, FORMAT_YYYY_MM_DD_HH_II_SS)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(timestamp)
	}
}

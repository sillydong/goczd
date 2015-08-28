package gotime
import "time"

const(
	FORMAT_YYYY_MM_DD="2006-01-02"
	FORMAT_YYYY_MM_DD_HH_II_SS="2006-01-02 15:04:05"
	Y="2016"
	M="01"
	D="02"
	H="15"
	I="04"
	S="05"
)

func GetTimestamp()(int64){
	return time.Now().Unix()
}

func GetTimeStr(format string)(string){
	return time.Now().Format(format)
}

func StrToTime(str,format string)(int64,error){
	timetime,err:=time.Parse(format,str)
	if err != nil {
		return 0,err
	}else{
		return timetime.Unix(),nil
	}
}

func TimeToStr(timestamp int64, format string)(string){
	return time.Unix(timestamp,0).Format(format)
}

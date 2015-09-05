package govalidate

import (
	"encoding/json"
	"github.com/cosiner/gohper/regexp"
	"net"
	"net/url"
	"strings"
)

//判断是否空字符串
func IsNull(str string) bool {
	return len(str) == 0
}

//判断是否邮箱
func IsEmail(data string) bool {
	return regexp.MustCompile("^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$").MatchString(data)
}

//是否ISBN
func IsISBN(str string, version int) bool {
	r, _ := regexp.Compile("[\\s-]+")
	sanitized := r.ReplaceAll([]byte(str), []byte(""))
	var checksum int32
	var i int32
	if version == 10 {
		if !regexp.MustCompile("^(?:[0-9]{9}X|[0-9]{10})$").MatchString(string(sanitized)) {
			return false
		}
		for i = 0; i < 9; i++ {
			checksum += (i + 1) * int32(sanitized[i]-'0')
		}
		if sanitized[9] == 'X' {
			checksum += 10 * 10
		} else {
			checksum += 10 * int32(sanitized[9]-'0')
		}
		if checksum%11 == 0 {
			return true
		}
		return false
	} else if version == 13 {
		if !regexp.MustCompile("^(?:[0-9]{13})$").MatchString(string(sanitized)) {
			return false
		}
		factor := []int32{1, 3}
		for i = 0; i < 12; i++ {
			checksum += factor[i%2] * int32(sanitized[i]-'0')
		}
		if (int32(sanitized[12]-'0'))-((10-(checksum%10))%10) == 0 {
			return true
		}
		return false
	}
	return IsISBN(str, 10) || IsISBN(str, 13)
}

//是否10位ISBN
func IsISBN10(str string) bool {
	return IsISBN(str, 10)
}

//是否13位ISBN
func IsISBN13(str string) bool {
	return IsISBN(str, 13)
}

//是否UUIDv3
func IsUUIDv3(str string) bool {
	return regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$").MatchString(str)
}

//是否UUIDv4.
func IsUUIDv4(str string) bool {
	return regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$").MatchString(str)
}

//是否UUIDv5
func IsUUIDv5(str string) bool {
	return regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$").MatchString(str)
}

//是否UUID
func IsUUID(str string) bool {
	return regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$").MatchString(str)
}

//是否字母
func IsAlpha(str string) bool {
	if IsNull(str) {
		return true
	}
	return regexp.MustCompile("^[a-zA-Z]+$").MatchString(str)
}

//是否字母数字组合
func IsAlphanumeric(str string) bool {
	if IsNull(str) {
		return true
	}
	return regexp.MustCompile("^[a-zA-Z0-9]+$").MatchString(str)
}

//是否数组
func IsNumeric(str string) bool {
	if IsNull(str) {
		return true
	}
	return regexp.MustCompile("^[-+]?[0-9]+$").MatchString(str)
}

//是否int
func IsInt(str string) bool {
	if IsNull(str) {
		return true
	}
	return regexp.MustCompile("^(?:[-+]?(?:0|[1-9][0-9]*))$").MatchString(str)
}

//是否float
func IsFloat(str string) bool {
	return regexp.MustCompile("^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$").MatchString(str)
}

//是否hex
func IsHexadecimal(str string) bool {
	return regexp.MustCompile("^[0-9a-fA-F]+$").MatchString(str)
}

//是否hex颜色值
func IsHexcolor(str string) bool {
	return regexp.MustCompile("^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$").MatchString(str)
}

//是否RGB颜色值
func IsRGBcolor(str string) bool {
	return regexp.MustCompile("^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$").MatchString(str)
}

//是否ASCII字符
func IsASCII(str string) bool {
	if IsNull(str) {
		return true
	}
	return regexp.MustCompile("^[\x00-\x7F]+$").MatchString(str)
}

//是否全角
func IsFullWidth(str string) bool {
	if IsNull(str) {
		return true
	}
	return regexp.MustCompile("[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]").MatchString(str)
}

//是否半角
func IsHalfWidth(str string) bool {
	if IsNull(str) {
		return true
	}
	return regexp.MustCompile("[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]").MatchString(str)
}

//是否BASE64
func IsBase64(str string) bool {
	return regexp.MustCompile("^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$").MatchString(str)
}

//是否可打印ASCII
func IsPrintableASCII(str string) bool {
	if IsNull(str) {
		return true
	}
	return regexp.MustCompile("^[\x20-\x7E]+$").MatchString(str)
}

//是否经度
func IsLatitude(str string) bool {
	return regexp.MustCompile("^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$").MatchString(str)
}

//是否纬度
func IsLongitude(str string) bool {
	return regexp.MustCompile("^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$").MatchString(str)
}

//是否URL链接
func IsURL(str string) bool {
	if str == "" || len(str) >= 2083 || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false
	}
	u, err := url.Parse(str)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	return regexp.MustCompile(`^((ftp|http|https):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|((([0-9a-zA-Z\.]*))|((www\.)?))?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?))(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`).MatchString(str)
}

//是否windows路径
func IsWinPath(str string) bool {
	return regexp.MustCompile(`^[a-zA-Z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*$`).MatchString(str)
}

//是否unix路径
func IsUnixPath(str string) bool {
	return regexp.MustCompile(`^((?:\/[a-zA-Z0-9\.\:]+(?:_[a-zA-Z0-9\:\.]+)*(?:\-[\:a-zA-Z0-9\.]+)*)+\/?)$`).MatchString(str)
}

//是否IP地址
func IsIP(str string) bool {
	return net.ParseIP(str) != nil
}

//是否IPV4
func IsIPv4(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && ip.To4() != nil
}

//是否IPV6
func IsIPv6(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && ip.To4() == nil
}

//是否MAC地址
func IsMAC(str string) bool {
	_, err := net.ParseMAC(str)
	return err == nil
}

//是否JSON
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

//判断是否手机号
func IsMobile(data string) bool {
	return regexp.MustCompile("^1[3875][0-9]{9}$").MatchString(data)
}

package godata

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
	"unsafe"
)

//string转[]byte
func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

//[]byte转string
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

const (
	ALPHABET     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	NUMERALS     = "0123456789"
	ALPHANUMERIC = NUMERALS + ALPHABET
)

//生成随机字符串
func RandomString(n int, randstring string) string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune(randstring)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//HMAC-SH1参数校验算法
func HMACSH1(params []string, sep, key string) string {
	sort.Strings(params)
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(strings.Join(params, sep)))
	return base64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))
}

//批量替换字符串
func StrReplace(s string, from, to []string, n int) string {
	if len(s) == 0 || len(from) == 0 || len(to) == 0 || len(from) != len(to) {
		return s
	}
	for key, valfrom := range from {
		valto := to[key]
		s = strings.Replace(s, valfrom, valto, n)
	}
	return s
}

//友好的byte显示
func FriendlyByte(b int64) string {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
	)

	switch {
	case b >= TB:
		return fmt.Sprintf("%.4fTB", float64(b)/float64(TB))
	case b >= GB:
		return fmt.Sprintf("%.4fGB", float64(b)/float64(GB))
	case b >= MB:
		return fmt.Sprintf("%.4fMB", float64(b)/float64(MB))
	case b >= KB:
		return fmt.Sprintf("%.4fKB", float64(b)/float64(KB))
	}
	return fmt.Sprintf("%.4fB", float64(b))
}

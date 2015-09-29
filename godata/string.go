package godata

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"math/rand"
	"sort"
	"strings"
	"time"
	"fmt"
)

//生成随机字符串
func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("AIJUVWXYZ012KLMNOPQRST3ijklmn4opqrstuvwxyz56BCDEFGH789abcdefghopqrstuvwxyz")
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
func FixByte(b int64)string{
	const (
		_ = iota
		KB = 1<<(10*iota)
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

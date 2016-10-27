package godata

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	fmt.Printf("%v\n", MD5([]byte("qqqqqq")))
}

func TestHash(t *testing.T) {
	fmt.Printf("%v\n", Hash([]byte("aaaaaaa")))
}

func TestCRC32(t *testing.T) {
	fmt.Printf("%v\n", CRC32([]byte("aaaaaaa")))
}

func TestSHA1(t *testing.T) {
	fmt.Printf("%v\n", SHA1([]byte("aaaaaaa")))
}

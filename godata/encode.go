package godata

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash/crc32"
	"hash/fnv"
)

func MD5(data []byte) string {
	hasher := md5.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

func Hash(data []byte) uint32 {
	hasher := fnv.New32()
	hasher.Write(data)
	return hasher.Sum32()
}

func CRC32(data []byte) uint32 {
	crc32q := crc32.MakeTable(0xD5828281)
	return crc32.Checksum(data, crc32q)
}

func SHA1(data []byte) string {
	hasher := sha1.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

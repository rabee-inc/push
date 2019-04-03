package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
	"unsafe"
)

// StrToMD5 ... 文字列のハッシュ(MD5)を取得する
func StrToMD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	hs := fmt.Sprintf("%x", h.Sum(nil))
	return strings.ToLower(hs)
}

// StrToSHA256 ... 文字列のハッシュ(SHA256)を取得する
func StrToSHA256(str string) string {
	c := sha256.Sum256([]byte(str))
	hs := hex.EncodeToString(c[:])
	return strings.ToLower(hs)
}

// StrToBytes ... 文字列をバイト列に変換する
func StrToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

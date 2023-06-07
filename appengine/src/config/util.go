package config

import (
	"crypto/md5"
	"fmt"
	"io"
)

func ToMD5(str string) string {
	h := md5.New()
	_, err := io.WriteString(h, str)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

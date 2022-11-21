package config

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/rabee-inc/go-pkg/deploy"
)

func GetFilePath(path string) string {
	if deploy.IsLocal() {
		return fmt.Sprintf("./%s", path)
	} else {
		return fmt.Sprintf("./appengine/push/%s", path)
	}
}

func ToMD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

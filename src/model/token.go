package model

import (
	"fmt"

	"github.com/aikizoku/push/src/lib/util"
)

// Token ...
type Token struct {
	Platform  string `firestore:""`
	DeviceID  string `firestore:""`
	Token     string `firestore:""`
	CreatedAt int64  `firestore:""`
}

// GenerateTokenDocID ...
func GenerateTokenDocID(pf string, deviceID string) string {
	return util.StrToMD5(fmt.Sprintf("%s_%s", pf, deviceID))
}

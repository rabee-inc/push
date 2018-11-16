package model

import (
	"fmt"

	"github.com/aikizoku/push/src/lib/util"
)

// PushTokenDatastore ... トークン(DataStore)
type PushTokenDatastore struct {
	ID        string `datastore:"-" boom:"id"`
	UserID    string ``
	Platform  string `datastore:",noindex"`
	DeviceID  string `datastore:",noindex"`
	Token     string `datastore:",noindex"`
	CreatedAt int64  `datastore:",noindex"`
}

// TokenFirestore ... トークン(FireStore)
type TokenFirestore struct {
	Platform  string `firestore:"platform"`
	DeviceID  string `firestore:"device_id"`
	Token     string `firestore:"token"`
	CreatedAt int64  `firestore:"created_at"`
}

// GeneratePushTokenKey ... Datastore用のKeyを作成する
func GeneratePushTokenKey(userID string, pf string, deviceID string) string {
	return util.StrToMD5(fmt.Sprintf("%s_%s_%s", userID, pf, deviceID))
}

// GenerateTokenDocID ... Firestore用のDocIDを作成する
func GenerateTokenDocID(pf string, deviceID string) string {
	return util.StrToMD5(fmt.Sprintf("%s_%s", pf, deviceID))
}

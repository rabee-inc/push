package model

import (
	"fmt"

	"github.com/rabee-inc/push/src/lib/util"
)

// PushToken ... トークン(DataStore)
type PushToken struct {
	ID        string `datastore:"-" boom:"id"`
	AppID     string `datastore:""`
	UserID    string `datastore:""`
	Platform  string `datastore:",noindex"`
	DeviceID  string `datastore:",noindex"`
	Token     string `datastore:",noindex"`
	CreatedAt int64  `datastore:",noindex"`
}

// GeneratePushTokenKey ... Datastore用のKeyを作成する
func GeneratePushTokenKey(appID string, userID string, pf string, deviceID string) string {
	return util.StrToMD5(fmt.Sprintf("%s::%s::%s::%s", appID, userID, pf, deviceID))
}

// TokenFirestore ... トークン(FireStore)
type TokenFirestore struct {
	Platform  string `firestore:"platform"`
	DeviceID  string `firestore:"device_id"`
	Token     string `firestore:"token"`
	CreatedAt int64  `firestore:"created_at"`
}

// GenerateTokenDocID ... Firestore用のDocIDを作成する
func GenerateTokenDocID(pf string, deviceID string) string {
	return util.StrToMD5(fmt.Sprintf("%s::%s", pf, deviceID))
}

// TokenMySQL ... トークン(MySQL)
type TokenMySQL struct {
	ID        string
	AppID     string
	UserID    string
	Platform  string
	DeviceID  string
	Token     string
	CreatedAt int64
	UpdatedAt int64
}

// GenerateTokenID ... MySQL用のIDを作成する
func GenerateTokenID(appID string, userID string, pf string, deviceID string) string {
	return util.StrToMD5(fmt.Sprintf("%s::%s::%s::%s", appID, userID, pf, deviceID))
}

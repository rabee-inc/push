package model

import (
	"fmt"

	"github.com/rabee-inc/push/src/lib/util"
)

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

// TableName ... TableNameを取得する
func (m *TokenMySQL) TableName() string {
	return "tokens"
}

// InsertColumns ... InsertColumnsを取得する
func (m *TokenMySQL) InsertColumns() []string {
	return []string{
		"app_id",
		"user_id",
		"platform_id",
		"token",
		"created_at",
		"updated_at",
	}
}

// GenerateTokenID ... MySQL用のIDを作成する
func GenerateTokenID(appID string, userID string, pf string, deviceID string) string {
	return util.StrToMD5(fmt.Sprintf("%s::%s::%s::%s", appID, userID, pf, deviceID))
}

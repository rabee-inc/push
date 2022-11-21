package model

import (
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/rabee-inc/push/appengine/push/src/config"
)

// Token ... トークン
type Token struct {
	ID        string                 `firestore:"-" cloudfirestore:"id"`
	Ref       *firestore.DocumentRef `firestore:"-" cloudfirestore:"ref"`
	Platform  string                 `firestore:"platform"`
	DeviceID  string                 `firestore:"device_id"`
	Token     string                 `firestore:"token"`
	CreatedAt int64                  `firestore:"created_at"`
}

// GenerateTokenDocID ... Firestore用のDocIDを作成する
func GenerateTokenDocID(pf string, deviceID string) string {
	return config.ToMD5(fmt.Sprintf("%s::%s", pf, deviceID))
}

// TokenRef ... コレクション参照を取得
func TokenRef(fCli *firestore.Client, appID string, userID string) *firestore.CollectionRef {
	return UserRef(fCli, appID).Doc(userID).Collection("tokens")
}

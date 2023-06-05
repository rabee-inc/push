package model

import (
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/rabee-inc/go-pkg/timeutil"
	"github.com/rabee-inc/push/appengine/push/src/config"
)

func TokenRef(cFirestore *firestore.Client, appID string, userID string) *firestore.CollectionRef {
	return UserRef(cFirestore, appID).Doc(userID).Collection("tokens")
}

type Token struct {
	ID        string                 `firestore:"-" cloudfirestore:"id"`
	Ref       *firestore.DocumentRef `firestore:"-" cloudfirestore:"ref"`
	Platform  string                 `firestore:"platform"`
	DeviceID  string                 `firestore:"device_id"`
	Token     string                 `firestore:"token"`
	CreatedAt int64                  `firestore:"created_at"`
}

func GenerateTokenDocID(platform string, deviceID string) string {
	return config.ToMD5(fmt.Sprintf("%s::%s", platform, deviceID))
}

func NewToken(
	id string,
	platform string,
	deviceID string,
	token string,
) *Token {
	now := timeutil.NowUnix()
	return &Token{
		ID:        id,
		Platform:  platform,
		DeviceID:  deviceID,
		Token:     token,
		CreatedAt: now,
	}
}

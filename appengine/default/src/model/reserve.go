package model

import (
	"cloud.google.com/go/firestore"

	"github.com/rabee-inc/push/appengine/default/src/config"
)

// Reserve ... 予約
type Reserve struct {
	ID         string                 `json:"id"          firestore:"-" cloudfirestore:"id"`
	Ref        *firestore.DocumentRef `json:"ref"         firestore:"-" cloudfirestore:"ref"`
	Message    *Message               `json:"message"     firestore:"message"`
	ReservedAt int64                  `json:"reserved_at" firestore:"reserved_at"`
	Status     config.ReserveStatus   `json:"status"      firestore:"status"`
	CreatedAt  int64                  `json:"created_at"  firestore:"created_at"`
	UpdatedAt  int64                  `json:"updated_at"  firestore:"updated_at"`
}

// ReserveRef ... コレクション参照を取得
func ReserveRef(fCli *firestore.Client, appID string) *firestore.CollectionRef {
	return fCli.Collection("push_apps").Doc(appID).Collection("reserves")
}

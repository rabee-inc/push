package model

import "cloud.google.com/go/firestore"

// UserRef ... コレクション参照を取得
func UserRef(fCli *firestore.Client, appID string) *firestore.CollectionRef {
	return AppRef(fCli).Doc(appID).Collection("users")
}

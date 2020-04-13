package model

import "cloud.google.com/go/firestore"

// AppRef ... コレクション参照を取得
func AppRef(fCli *firestore.Client) *firestore.CollectionRef {
	return fCli.Collection("push_apps")
}

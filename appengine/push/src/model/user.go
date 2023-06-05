package model

import "cloud.google.com/go/firestore"

func UserRef(cFirestore *firestore.Client, appID string) *firestore.CollectionRef {
	return AppRef(cFirestore).Doc(appID).Collection("users")
}

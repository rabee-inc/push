package model

import "cloud.google.com/go/firestore"

func AppRef(cFirestore *firestore.Client) *firestore.CollectionRef {
	return cFirestore.Collection("push_apps")
}

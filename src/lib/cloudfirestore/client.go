package cloudfirestore

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/appengine/log"
)

// NewClient ... Firestoreのクライアントを取得する
func NewClient(ctx context.Context) (*firestore.Client, error) {
	pID := os.Getenv("PROJECT_ID")
	if pID == "" {
		err := fmt.Errorf("failed to get env project_id")
		panic(err)
	}
	cfg := &firebase.Config{ProjectID: pID}
	app, err := firebase.NewApp(ctx, cfg)
	if err != nil {
		log.Errorf(ctx, "firebase.NewApp error: %s", err.Error())
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Errorf(ctx, "app.Firestore error: %s", err.Error())
		return nil, err
	}
	return client, nil
}

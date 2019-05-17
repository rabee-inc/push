package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	_ "github.com/rabee-inc/push/src/statik" // バイナリ化したファイル
	"github.com/rakyll/statik/fs"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

// FCMApps ... FCMのApps
type FCMApps struct {
	AppIDs []string `json:"app_ids"`
}

// FCMEnv ... FCMのEnv
type FCMEnv struct {
	AppID     string
	ServerKey string `json:"server_key"`
}

// GetFCMAppIDs ... AppIDsを取得する
func GetFCMAppIDs() []string {
	jsonData := getFile("/apps.json")
	var fcmApps FCMApps
	err := json.Unmarshal(jsonData, &fcmApps)
	if err != nil {
		panic(err)
	}
	return fcmApps.AppIDs
}

// GetClient ... Clientを取得する
func GetClient(e string, appID string) *messaging.Client {
	ctx := context.Background()
	jsonData := getFile(fmt.Sprintf("/credentials/%s/%s.json", e, appID))
	cre, err := google.CredentialsFromJSON(ctx, jsonData)
	if err != nil {
		panic(err)
	}
	opt := option.WithCredentials(cre)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
	cli, err := app.Messaging(ctx)
	if err != nil {
		panic(err)
	}
	return cli
}

// GetFCMEnv ... Envを取得する
func GetFCMEnv(e string, appID string) *FCMEnv {
	jsonData := getFile(fmt.Sprintf("/env/%s/%s.json", e, appID))
	var fcmEnv FCMEnv
	err := json.Unmarshal(jsonData, &fcmEnv)
	if err != nil {
		panic(err)
	}
	fcmEnv.AppID = appID
	return &fcmEnv
}

func getFile(path string) []byte {
	fileSystem, err := fs.New()
	if err != nil {
		panic(err)
	}

	file, err := fileSystem.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return b
}

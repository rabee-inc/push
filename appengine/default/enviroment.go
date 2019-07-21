package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Environment ... 環境変数
type Environment struct {
	Port              int    `envconfig:"PORT"                           default:"8080"`
	Deploy            string `envconfig:"DEPLOY"                         required:"true"`
	ProjectID         string `envconfig:"PROJECT_ID"                     required:"true"`
	LocationID        string `envconfig:"LOCATION_ID"                    default:"asia-northeast1"`
	ServiceID         string `envconfig:"SERVICE_ID"                     required:"true"`
	CredentialsPath   string `envconfig:"GOOGLE_APPLICATION_CREDENTIALS" required:"true"`
	MinLogSeverity    string `envconfig:"MIN_LOG_SEVERITY"               required:"true"`
	InternalAuthToken string `envconfig:"INTERNAL_AUTH_TOKEN"            required:"true"`
	MySQLPushHost     string `envconfig:"MYSQL_PUSH_HOST"`
	MySQLPushUser     string `envconfig:"MYSQL_PUSH_USER"`
	MySQLPushPassword string `envconfig:"MYSQL_PUSH_PASSWORD"`
	MySQLPushDB       string `envconfig:"MYSQL_PUSH_DB"`
}

// Get ... 環境変数を取得する
func (e *Environment) Get() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	err = envconfig.Process("", e)
	if err != nil {
		panic(err)
	}
}

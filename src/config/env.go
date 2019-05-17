package config

import (
	"os"
)

// GetEnv ... 現在の環境を取得する
func GetEnv() string {
	e := os.Getenv("ENV")
	if e == "" {
		panic("no config ENV")
	}
	return e
}

// IsEnvLocal ... 現在の環境がローカルか判定する
func IsEnvLocal() bool {
	return GetEnv() == "local"
}

// IsEnvStaging ... 現在の環境がステージングか判定する
func IsEnvStaging() bool {
	return GetEnv() == "staging"
}

// IsEnvProduction ... 現在の環境が本番か判定する
func IsEnvProduction() bool {
	return GetEnv() == "production"
}

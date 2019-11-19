package service

import (
	"context"
)

// Register ... プッシュ通知登録に関するサービス
type Register interface {
	SetToken(ctx context.Context, appID string, userID string, platform string, deviceID string, token string) error
	DeleteToken(ctx context.Context, appID string, userID string, platform string, deviceID string) error
}

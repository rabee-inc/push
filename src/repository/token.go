package repository

import (
	"context"
)

// Token ... トークン
type Token interface {
	GetListByUserID(ctx context.Context, appID string, userID string) ([]string, error)
	Put(ctx context.Context, appID string, userID string, platform string, deviceID string, token string) error
	Delete(ctx context.Context, appID string, userID string, platform string, deviceID string) error
}

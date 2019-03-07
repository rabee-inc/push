package repository

import (
	"context"
)

// Token ... トークン
type Token interface {
	GetListByUserID(ctx context.Context, userID string) ([]string, error)
	Put(ctx context.Context, userID string, platform string, deviceID string, token string) error
	Delete(ctx context.Context, userID string, platform string, deviceID string) error
}

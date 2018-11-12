package repository

import (
	"context"
)

// Token ... トークン
type Token interface {
	Put(ctx context.Context, userID string, platform string, deviceID string, token string) error
}

package service

import (
	"context"
)

// Entry ... 通知を登録する
type Entry interface {
	Token(ctx context.Context, userID string, platform string, deviceID string, token string) error
}

package service

import (
	"context"
)

// Entry ...
type Entry interface {
	Token(ctx context.Context, userID string, platform string, deviceID string, token string) error
}

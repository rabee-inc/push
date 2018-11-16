package service

import (
	"context"

	"github.com/aikizoku/push/src/model"
)

// Sender ... 通知を送信する
type Sender interface {
	SendMessageToUserIDs(ctx context.Context, userIDs []string, msg *model.Message) error
	SendMessageToUserID(ctx context.Context, userID string, msg *model.Message) error
	SendMessageToToken(ctx context.Context, token string, msg *model.Message) error
}

package service

import (
	"context"

	"github.com/aikizoku/push/src/model"
)

// Sender ... 通知を送信する
type Sender interface {
	SendMessageByUserIDs(ctx context.Context, userIDs []string, msg *model.Message) error
	SendMessageByUserID(ctx context.Context, userID string, msg *model.Message) error
	SendMessageByToken(ctx context.Context, token string, msg *model.Message) error
}

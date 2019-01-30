package service

import (
	"context"

	"github.com/rabee-inc/push/src/model"
)

// Sender ... 通知を送信する
type Sender interface {
	MessageByUserIDs(ctx context.Context, userIDs []string, msg *model.Message) error
	MessageByUserID(ctx context.Context, userID string, msg *model.Message) error
	MessageByToken(ctx context.Context, token string, msg *model.Message) error
}

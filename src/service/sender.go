package service

import (
	"context"

	"github.com/rabee-inc/push/src/model"
)

// Sender ... 通知を送信する
type Sender interface {
	MessageByUserIDs(ctx context.Context, appID string, userIDs []string, msg *model.Message) error
	MessageByUserID(ctx context.Context, appID string, userID string, msg *model.Message) error
	MessageByToken(ctx context.Context, appID string, token string, msg *model.Message) error
}

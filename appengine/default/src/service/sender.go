package service

import (
	"context"

	"github.com/rabee-inc/push/appengine/default/src/model"
)

// Sender ... 通知を送信する
type Sender interface {
	MessageByAll(ctx context.Context, appID string, msg *model.Message) error
	MessageByUserIDs(ctx context.Context, appID string, userIDs []string, msg *model.Message) error
	MessageByUserID(ctx context.Context, appID string, userID string, msg *model.Message) error
}

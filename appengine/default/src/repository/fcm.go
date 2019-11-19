package repository

import (
	"context"

	"github.com/rabee-inc/push/appengine/default/src/model"
)

// Fcm ... FCMに関するリポジトリ
type Fcm interface {
	SubscribeTopic(ctx context.Context, appID string, topic string, tokens []string) error
	UnsubscribeTopic(ctx context.Context, appID string, topic string, tokens []string) error
	SendMessage(ctx context.Context, appID string, token string, src *model.Message) error
}

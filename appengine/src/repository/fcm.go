package repository

import (
	"context"

	"github.com/rabee-inc/push/appengine/src/model"
)

// Fcm ... FCMに関するリポジトリ
type Fcm interface {
	SendMessage(ctx context.Context, appID string, token string, src *model.Message) error
}

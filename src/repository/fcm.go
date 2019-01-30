package repository

import (
	"context"

	"github.com/rabee-inc/push/src/model"
)

// Fcm ... FCM
type Fcm interface {
	SendMessage(ctx context.Context, token string, src *model.Message) error
}

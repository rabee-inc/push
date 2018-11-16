package repository

import (
	"context"

	"github.com/aikizoku/push/src/model"
)

// Fcm ... FCM
type Fcm interface {
	SendMessage(ctx context.Context, token string, src *model.Message) error
}

package service

import (
	"context"

	"github.com/rabee-inc/push/appengine/default/src/model"
)

// Sender ... 通知を送信する
type Sender interface {
	AllUsers(
		ctx context.Context,
		appID string,
		msg *model.Message) error
	Users(
		ctx context.Context,
		appID string,
		userIDs []string,
		msg *model.Message) error
	User(
		ctx context.Context,
		appID string,
		userID string,
		msg *model.Message) error
	Reserved(
		ctx context.Context,
		appID string) error
}

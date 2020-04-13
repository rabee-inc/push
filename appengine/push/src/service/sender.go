package service

import (
	"context"

	"github.com/rabee-inc/push/appengine/push/src/model"
)

// Sender ... 通知を送信する
type Sender interface {
	AllUsers(
		ctx context.Context,
		appID string,
		pushID string,
		msg *model.Message) error
	Users(
		ctx context.Context,
		appID string,
		userIDs []string,
		pushID string,
		msg *model.Message) error
	User(
		ctx context.Context,
		appID string,
		userID string,
		pushID string,
		msg *model.Message) error
	Reserved(
		ctx context.Context,
		appID string) error
}

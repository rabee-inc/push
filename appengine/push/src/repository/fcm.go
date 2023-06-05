package repository

import (
	"context"

	"github.com/rabee-inc/push/appengine/push/src/model"
)

type FCM interface {
	SubscribeTopic(
		ctx context.Context,
		appID string,
		topic string,
		tokens []*model.Token,
	) error

	UnsubscribeTopic(
		ctx context.Context,
		appID string,
		topic string,
		tokens []*model.Token,
	) error

	SendMessageByTokens(
		ctx context.Context,
		appID string,
		tokens []*model.Token,
		pushID string,
		message *model.Message,
	) error

	SendMessageByTopic(
		ctx context.Context,
		appID string,
		topic string,
		pushID string,
		message *model.Message,
	) error
}

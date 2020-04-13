package repository

import (
	"context"

	"github.com/rabee-inc/push/appengine/push/src/model"
)

// Fcm ... FCMに関するリポジトリ
type Fcm interface {
	SubscribeTopic(
		ctx context.Context,
		appID string,
		topic string,
		tokens []string) error
	UnsubscribeTopic(
		ctx context.Context,
		appID string,
		topic string,
		tokens []string) error
	SendMessageByTokens(
		ctx context.Context,
		appID string,
		tokens []string,
		pushID string,
		src *model.Message) error
	SendMessageByTopic(
		ctx context.Context,
		appID string,
		topic string,
		pushID string,
		src *model.Message) error
}

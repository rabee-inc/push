package repository

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/messaging"
	"github.com/rabee-inc/go-pkg/log"
	"github.com/rabee-inc/go-pkg/sliceutil"
	"github.com/rabee-inc/push/appengine/src/model"
)

type fcm struct {
	cMessaging *messaging.Client
	serverKey  string
}

func NewFCM(
	cMessaging *messaging.Client,
	serverKey string,
) FCM {
	return &fcm{
		cMessaging,
		serverKey,
	}
}

func (r *fcm) SubscribeTopic(
	ctx context.Context,
	appID string,
	topic string,
	tokens []*model.Token,
) error {
	ts := r.filterMapToken(tokens)
	if len(ts) == 0 {
		return nil
	}
	res, err := r.cMessaging.SubscribeToTopic(ctx, ts, topic)
	if err != nil {
		log.Error(ctx, err)
		return err
	}
	if res.FailureCount > 0 {
		for _, rErr := range res.Errors {
			return log.Warninge(ctx, "SubscribeToTopic index: %d, reason: %s", rErr.Index, rErr.Reason)
		}
	}
	return nil
}

func (r *fcm) UnsubscribeTopic(
	ctx context.Context,
	appID string,
	topic string,
	tokens []*model.Token,
) error {
	ts := r.filterMapToken(tokens)
	if len(ts) == 0 {
		return nil
	}
	res, err := r.cMessaging.UnsubscribeFromTopic(ctx, ts, topic)
	if err != nil {
		log.Error(ctx, err)
		return err
	}
	if res.FailureCount > 0 {
		for _, rErr := range res.Errors {
			return log.Warninge(ctx, "UnsubscribeFromTopic index: %d, reason: %s", rErr.Index, rErr.Reason)
		}
	}
	return nil
}

func (r *fcm) SendMessageByTokens(
	ctx context.Context,
	appID string,
	tokens []*model.Token,
	pushID string,
	message *model.Message,
) error {
	ts := r.filterMapToken(tokens)
	if len(ts) == 0 {
		return nil
	}
	msg := r.generateMessage(pushID, message)
	multiMsg := &messaging.MulticastMessage{
		Tokens:       ts,
		Notification: msg.Notification,
		Data:         msg.Data,
		APNS:         msg.APNS,
		Android:      msg.Android,
		Webpush:      msg.Webpush,
	}
	multiRes, err := r.cMessaging.SendEachForMulticast(ctx, multiMsg)
	if err != nil {
		log.Warning(ctx, err)
		return err
	}
	if multiRes.FailureCount > 0 {
		for _, res := range multiRes.Responses {
			if res == nil || res.Error == nil {
				continue
			}
			log.Warning(ctx, res.Error)
			// 個別の送信エラーは無視
		}
	}
	return nil
}

func (r *fcm) SendMessageByTopic(
	ctx context.Context,
	appID string,
	topic string,
	pushID string,
	message *model.Message,
) error {
	msg := r.generateMessage(pushID, message)
	msg.Topic = topic
	_, err := r.cMessaging.Send(ctx, msg)
	if err != nil {
		log.Warning(ctx, err)
		return err
	}
	return nil
}

func (r *fcm) filterMapToken(tokens []*model.Token) []string {
	return sliceutil.FilterMap(tokens, func(token *model.Token) (bool, string) {
		return token.Token != "", token.Token
	})
}

func (r *fcm) generateMessage(
	pushID string,
	message *model.Message,
) *messaging.Message {
	if message.IOS == nil {
		message.IOS = &model.MessageIOS{
			Badge: 1,
		}
	}
	if message.Android == nil {
		message.Android = &model.MessageAndroid{}
	}
	if message.Web == nil {
		message.Web = &model.MessageWeb{}
	}

	msg := &messaging.Message{
		Notification: &messaging.Notification{
			Title: message.Title,
			Body:  message.Body,
		},
		Data: message.Data,
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Badge:          &message.IOS.Badge,
					Sound:          message.IOS.Sound,
					MutableContent: true,
				},
				CustomData: map[string]any{
					"push_id":                 pushID,
					"notification_foreground": true,
				},
			},
			FCMOptions: &messaging.APNSFCMOptions{
				ImageURL: message.IOS.ImageURL,
			},
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				ClickAction: message.Android.ClickAction,
				Sound:       message.Android.Sound,
				Tag:         message.Android.Tag,
			},
			Data: map[string]string{
				"push_id":                 pushID,
				"notification_foreground": "true",
			},
		},
		Webpush: &messaging.WebpushConfig{
			Headers: map[string]string{
				"Content-Type":  "application/json",
				"Authorization": fmt.Sprintf("Bearer %s", r.serverKey),
			},
			Notification: &messaging.WebpushNotification{
				Icon: message.Web.Icon,
			},
		},
	}
	return msg
}

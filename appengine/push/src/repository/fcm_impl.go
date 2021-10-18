package repository

import (
	"context"
	"fmt"

	"firebase.google.com/go/messaging"

	"github.com/rabee-inc/go-pkg/log"
	"github.com/rabee-inc/push/appengine/push/src/model"
)

type fcm struct {
	fCli   *messaging.Client
	svrKey string
}

func (r *fcm) SubscribeTopic(ctx context.Context, appID string, topic string, tokens []string) error {
	res, err := r.fCli.SubscribeToTopic(ctx, tokens, topic)
	if err != nil {
		log.Errorm(ctx, "r.fCli.SubscribeToTopic", err)
		return err
	}
	if res.FailureCount > 0 {
		for _, rerr := range res.Errors {
			err = log.Warninge(ctx, "SubscribeToTopic index: %d, reason: %s", rerr.Index, rerr.Reason)
			return err
		}
	}
	return nil
}

func (r *fcm) UnsubscribeTopic(ctx context.Context, appID string, topic string, tokens []string) error {
	res, err := r.fCli.UnsubscribeFromTopic(ctx, tokens, topic)
	if err != nil {
		log.Errorm(ctx, "r.fCli.UnsubscribeFromTopic", err)
		return err
	}
	if res.FailureCount > 0 {
		for _, rerr := range res.Errors {
			err = log.Warninge(ctx, "UnsubscribeFromTopic index: %d, reason: %s", rerr.Index, rerr.Reason)
			return err
		}
	}
	return nil
}

func (r *fcm) SendMessageByTokens(ctx context.Context, appID string, tokens []string, pushID string, src *model.Message) error {
	msg := r.generateMessage(ctx, appID, pushID, src)
	mmsg := &messaging.MulticastMessage{
		Tokens:       tokens,
		Notification: msg.Notification,
		Data:         msg.Data,
		APNS:         msg.APNS,
		Android:      msg.Android,
		Webpush:      msg.Webpush,
	}
	_, err := r.fCli.SendMulticast(ctx, mmsg)
	if err != nil {
		log.Warningm(ctx, "r.fCli.SendMulticast", err)
		return err
	}
	return nil
}

func (r *fcm) SendMessageByTopic(ctx context.Context, appID string, topic string, pushID string, src *model.Message) error {
	msg := r.generateMessage(ctx, appID, pushID, src)
	msg.Topic = topic
	_, err := r.fCli.Send(ctx, msg)
	if err != nil {
		log.Warningm(ctx, "r.fCli.Send", err)
		return err
	}
	return nil
}

func (r *fcm) generateMessage(ctx context.Context, appID string, pushID string, src *model.Message) *messaging.Message {
	if src.IOS == nil {
		src.IOS = &model.MessageIOS{
			Badge: 1,
		}
	}
	if src.Android == nil {
		src.Android = &model.MessageAndroid{}
	}
	if src.Web == nil {
		src.Web = &model.MessageWeb{}
	}

	msg := &messaging.Message{
		Notification: &messaging.Notification{
			Title: src.Title,
			Body:  src.Body,
		},
		Data: src.Data,
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Badge:          &src.IOS.Badge,
					Sound:          src.IOS.Sound,
					MutableContent: true,
				},
				CustomData: map[string]interface{}{
					"push_id":                 pushID,
					"notification_foreground": true,
				},
			},
			FCMOptions: &messaging.APNSFCMOptions{
				ImageURL: src.IOS.ImageURL,
			},
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				ClickAction: src.Android.ClickAction,
				Sound:       src.Android.Sound,
				Tag:         src.Android.Tag,
			},
			Data: map[string]string{
				"push_id":                 pushID,
				"notification_foreground": "true",
			},
		},
		Webpush: &messaging.WebpushConfig{
			Headers: map[string]string{
				"Content-Type":  "application/json",
				"Authorization": fmt.Sprintf("Bearer %s", r.svrKey),
			},
			Notification: &messaging.WebpushNotification{
				Icon: src.Web.Icon,
			},
		},
	}
	return msg
}

// NewFcm ... リポジトリを作成する
func NewFcm(fCli *messaging.Client, svrKey string) Fcm {
	return &fcm{
		fCli:   fCli,
		svrKey: svrKey,
	}
}

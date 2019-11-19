package repository

import (
	"context"
	"fmt"

	"firebase.google.com/go/messaging"

	"github.com/rabee-inc/push/appengine/default/src/lib/log"
	"github.com/rabee-inc/push/appengine/default/src/model"
)

type fcm struct {
	fClis   map[string]*messaging.Client
	svrKeys map[string]string
}

func (r *fcm) SubscribeTopic(ctx context.Context, appID string, topic string, tokens []string) error {
	res, err := r.fClis[appID].SubscribeToTopic(ctx, tokens, topic)
	if err != nil {
		log.Errorm(ctx, "r.fClis[appID].SubscribeToTopic", err)
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
	res, err := r.fClis[appID].UnsubscribeFromTopic(ctx, tokens, topic)
	if err != nil {
		log.Errorm(ctx, "r.fClis[appID].UnsubscribeFromTopic", err)
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

func (r *fcm) SendMessageByTokens(ctx context.Context, appID string, tokens []string, src *model.Message) error {
	msg := r.generateMessage(ctx, appID, src)
	mmsg := &messaging.MulticastMessage{
		Tokens:       tokens,
		Notification: msg.Notification,
		Data:         msg.Data,
		APNS:         msg.APNS,
		Android:      msg.Android,
		Webpush:      msg.Webpush,
	}
	_, err := r.fClis[appID].SendMulticast(ctx, mmsg)
	if err != nil {
		log.Warningm(ctx, "r.fClis[appID].SendMulticast", err)
		return err
	}
	return nil
}

func (r *fcm) SendMessageByTopic(ctx context.Context, appID string, topic string, src *model.Message) error {
	msg := r.generateMessage(ctx, appID, src)
	msg.Topic = topic
	_, err := r.fClis[appID].Send(ctx, msg)
	if err != nil {
		log.Warningm(ctx, "r.fClis[appID].Send", err)
		return err
	}
	return nil
}

func (r *fcm) generateMessage(ctx context.Context, appID string, src *model.Message) *messaging.Message {
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
					Badge: &src.IOS.Badge,
					Sound: src.IOS.Sound,
				},
			},
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				ClickAction: src.Android.ClickAction,
				Sound:       src.Android.Sound,
				Tag:         src.Android.Tag,
			},
		},
		Webpush: &messaging.WebpushConfig{
			Headers: map[string]string{
				"Content-Type":  "application/json",
				"Authorization": fmt.Sprintf("Bearer %s", r.svrKeys[appID]),
			},
			Notification: &messaging.WebpushNotification{
				Icon: src.Web.Icon,
			},
		},
	}
	return msg
}

// NewFcm ... リポジトリを作成する
func NewFcm(fClis map[string]*messaging.Client, svrKeys map[string]string) Fcm {
	return &fcm{
		fClis:   fClis,
		svrKeys: svrKeys,
	}
}

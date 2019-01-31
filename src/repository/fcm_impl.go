package repository

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/rabee-inc/push/src/lib/log"
	"github.com/rabee-inc/push/src/model"
)

type fcm struct {
}

// SendMessage ... FCMにプッシュ通知送信を登録する
func (r *fcm) SendMessage(ctx context.Context, token string, src *model.Message) error {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Errorm(ctx, "firebase.NewApp", err)
		return err
	}

	cli, err := app.Messaging(ctx)
	if err != nil {
		log.Errorm(ctx, "app.Messaging", err)
		return err
	}

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
		Token: token,
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
			Notification: &messaging.WebpushNotification{
				Icon: src.Web.Icon,
			},
		},
	}

	_, err = cli.Send(ctx, msg)
	if err != nil {
		log.Errorm(ctx, "cli.Send", err)
		return err
	}
	return nil
}

// NewFcm ... リポジトリを作成する
func NewFcm() Fcm {
	return &fcm{}
}

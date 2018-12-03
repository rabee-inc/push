package repository

import (
	"context"

	"github.com/aikizoku/push/src/config"
	"github.com/aikizoku/push/src/lib/log"
	"github.com/aikizoku/push/src/lib/util"
	"github.com/aikizoku/push/src/model"
	_ "go.mercari.io/datastore/aedatastore" // mercari/datastoreの初期化
	"go.mercari.io/datastore/boom"
)

type tokenDatastore struct {
}

func (r *tokenDatastore) GetMultiToUserID(ctx context.Context, userID string) ([]string, error) {
	tokens := []string{}
	ret := []*model.PushTokenDatastore{}
	b, err := boom.FromContext(ctx)
	if err != nil {
		log.Errorf(ctx, "boom from context error: %s", err.Error())
		return tokens, err
	}
	q := b.NewQuery(config.KindPushToken).Filter("UserID =", userID)
	_, err = b.GetAll(q, &ret)
	if err != nil {
		log.Errorf(ctx, "get by query error: "+err.Error())
		return tokens, err
	}
	for _, r := range ret {
		tokens = append(tokens, r.Token)
	}
	return tokens, nil
}

func (r *tokenDatastore) Put(ctx context.Context, userID string, platform string, deviceID string, token string) error {
	src := &model.PushTokenDatastore{
		UserID:    userID,
		Platform:  platform,
		DeviceID:  deviceID,
		Token:     token,
		CreatedAt: util.TimeNowUnix(),
	}
	b, err := boom.FromContext(ctx)
	if err != nil {
		log.Errorf(ctx, "boom from context error: %s", err.Error())
		return err
	}
	_, err = b.Put(src)
	if err != nil {
		log.Errorf(ctx, "b.Put error: %s", err.Error())
		return err
	}
	return nil
}

// NewTokenDatastore ... TokenDatastoreを作成する
func NewTokenDatastore() Token {
	return &tokenDatastore{}
}

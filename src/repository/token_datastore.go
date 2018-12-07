package repository

import (
	"context"

	"github.com/aikizoku/push/src/config"
	"github.com/aikizoku/push/src/lib/log"
	"github.com/aikizoku/push/src/lib/util"
	"github.com/aikizoku/push/src/model"
	"go.mercari.io/datastore"
	_ "go.mercari.io/datastore/aedatastore" // mercari/datastoreの初期化
	"go.mercari.io/datastore/boom"
)

type tokenDatastore struct {
}

func (r *tokenDatastore) GetListByUserID(ctx context.Context, userID string) ([]string, error) {
	b, err := boom.FromContext(ctx)
	if err != nil {
		log.Errorm(ctx, "boom.FromContext", err)
		return []string{}, err
	}
	q := b.NewQuery(config.KindPushToken).Filter("UserID =", userID).KeysOnly()
	keys, err := b.GetAll(q, nil)
	if err != nil {
		log.Errorf(ctx, "b.GetAll error: %s", err.Error())
		return []string{}, err
	}
	ids := []string{}
	for _, key := range keys {
		ids = append(ids, key.Name())
	}
	tokens, err := r.getMulti(ctx, ids)
	if err != nil {
		log.Errorf(ctx, "r.getMulti error: %s", err.Error())
		return []string{}, err
	}
	return tokens, nil
}

func (r *tokenDatastore) getMulti(ctx context.Context, ids []string) ([]string, error) {
	tokens := []string{}
	b, err := boom.FromContext(ctx)
	if err != nil {
		log.Errorm(ctx, "boom.FromContext", err)
		return tokens, err
	}
	bt := b.Batch()
	for _, id := range ids {
		dst := &model.PushTokenDatastore{
			ID: id,
		}
		bt.Get(dst, func(err error) error {
			if err != nil {
				if err == datastore.ErrNoSuchEntity {
					return nil
				}
				log.Errorf(ctx, "bt.Get error: %s, id: %s", err.Error(), dst.ID)
				return err
			}
			tokens = append(tokens, dst.Token)
			return nil
		})
	}
	err = bt.Exec()
	if err != nil {
		log.Errorf(ctx, "bt.Exec error: %s", err.Error())
		return tokens, err
	}
	return tokens, nil
}

func (r *tokenDatastore) Put(ctx context.Context, userID string, platform string, deviceID string, token string) error {
	id := model.GeneratePushTokenKey(userID, platform, deviceID)
	src := &model.PushTokenDatastore{
		ID:        id,
		UserID:    userID,
		Platform:  platform,
		DeviceID:  deviceID,
		Token:     token,
		CreatedAt: util.TimeNowUnix(),
	}
	b, err := boom.FromContext(ctx)
	if err != nil {
		log.Errorm(ctx, "boom.FromContext", err)
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

package repository

import (
	"context"

	"github.com/rabee-inc/push/src/config"
	"github.com/rabee-inc/push/src/lib/log"
	"github.com/rabee-inc/push/src/lib/util"
	"github.com/rabee-inc/push/src/model"
	"go.mercari.io/datastore"
	_ "go.mercari.io/datastore/aedatastore" // mercari/datastoreの初期化
	"go.mercari.io/datastore/boom"
)

type tokenDatastore struct {
}

// GetListByUserID ... ユーザーIDに紐づくトークンリストを取得する
func (r *tokenDatastore) GetListByUserID(ctx context.Context, userID string) ([]string, error) {
	b, err := boom.FromContext(ctx)
	if err != nil {
		log.Errorm(ctx, "boom.FromContext", err)
		return []string{}, err
	}
	q := b.NewQuery(config.KindPushToken).Filter("UserID =", userID).KeysOnly()
	keys, err := b.GetAll(q, nil)
	if err != nil {
		log.Errorm(ctx, "b.GetAll", err)
		return []string{}, err
	}
	ids := []string{}
	for _, key := range keys {
		ids = append(ids, key.Name())
	}
	tokens, err := r.getMulti(ctx, ids)
	if err != nil {
		log.Errorm(ctx, "r.getMulti", err)
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
				log.Errorm(ctx, "bt.Get", err)
				return err
			}
			tokens = append(tokens, dst.Token)
			return nil
		})
	}
	err = bt.Exec()
	if err != nil {
		log.Errorm(ctx, "bt.Exec", err)
		return tokens, err
	}
	return tokens, nil
}

// Put ... トークンを登録する
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
		log.Errorm(ctx, "b.Put", err)
		return err
	}
	return nil
}

// Delete ... トークンを削除する
func (r *tokenDatastore) Delete(ctx context.Context, userID string, platform string, deviceID string) error {
	id := model.GeneratePushTokenKey(userID, platform, deviceID)
	src := &model.PushTokenDatastore{
		ID: id,
	}
	b, err := boom.FromContext(ctx)
	if err != nil {
		log.Errorm(ctx, "boom.FromContext", err)
		return err
	}
	err = b.Delete(src)
	if err != nil {
		log.Errorm(ctx, "b.Delete", err)
		return err
	}
	return nil
}

// NewTokenDatastore ... リポジトリを作成する
func NewTokenDatastore() Token {
	return &tokenDatastore{}
}

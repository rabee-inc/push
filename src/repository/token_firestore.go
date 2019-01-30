package repository

import (
	"context"

	"github.com/rabee-inc/push/src/config"
	"github.com/rabee-inc/push/src/lib/cloudfirestore"
	"github.com/rabee-inc/push/src/lib/log"
	"github.com/rabee-inc/push/src/lib/util"
	"github.com/rabee-inc/push/src/model"
	"google.golang.org/api/iterator"
)

type tokenFirestore struct {
}

// GetListByUserID ... ユーザーIDに紐づくトークンリストを取得する
func (r *tokenFirestore) GetListByUserID(ctx context.Context, userID string) ([]string, error) {
	var tokens []string
	cli, err := cloudfirestore.NewClient(ctx)
	if err != nil {
		log.Errorm(ctx, "cloudfirestore.NewClient", err)
		return tokens, err
	}
	iter := cli.Collection(config.CollectionUsers).Doc(userID).Collection(config.CollectionTokens).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Errorm(ctx, "iter.Next", err)
			return tokens, err
		}
		token := &model.TokenFirestore{}
		err = doc.DataTo(token)
		if err != nil {
			log.Errorm(ctx, "doc.DataTo", err)
			return nil, err
		}
		tokens = append(tokens, token.Token)
	}
	return tokens, nil
}

// Put ... トークンを登録する
func (r *tokenFirestore) Put(ctx context.Context, userID string, platform string, deviceID string, token string) error {
	docID := model.GenerateTokenDocID(platform, deviceID)
	src := &model.TokenFirestore{
		Platform:  platform,
		DeviceID:  deviceID,
		Token:     token,
		CreatedAt: util.TimeNowUnix(),
	}
	cli, err := cloudfirestore.NewClient(ctx)
	if err != nil {
		log.Errorm(ctx, "cloudfirestore.NewClient", err)
		return err
	}
	ret, err := cli.Collection(config.CollectionUsers).Doc(userID).Collection(config.CollectionTokens).Doc(docID).Set(ctx, src)
	if err != nil {
		log.Errorm(ctx, "cli.Set", err)
		return err
	}
	log.Debugf(ctx, "UpdateTime: %s", ret.UpdateTime)
	return nil
}

// NewTokenFirestore ... リポジトリを作成する
func NewTokenFirestore() Token {
	return &tokenFirestore{}
}

package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/rabee-inc/push/src/config"
	"github.com/rabee-inc/push/src/lib/log"
	"github.com/rabee-inc/push/src/lib/util"
	"github.com/rabee-inc/push/src/model"
	"google.golang.org/api/iterator"
)

type tokenFirestore struct {
	client *firestore.Client
}

// GetListByUserID ... ユーザーIDに紐づくトークンリストを取得する
func (r *tokenFirestore) GetListByUserID(ctx context.Context, userID string) ([]string, error) {
	var tokens []string
	iter := r.client.
		Collection(config.CollectionUsers).
		Doc(userID).
		Collection(config.CollectionTokens).
		Documents(ctx)
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
	_, err := r.client.
		Collection(config.CollectionUsers).
		Doc(userID).
		Collection(config.CollectionTokens).
		Doc(docID).
		Set(ctx, src)
	if err != nil {
		log.Errorm(ctx, "r.client.Set", err)
		return err
	}
	return nil
}

// Delete ... トークンを削除する
func (r *tokenFirestore) Delete(ctx context.Context, userID string, platform string, deviceID string) error {
	docID := model.GenerateTokenDocID(platform, deviceID)
	_, err := r.client.
		Collection(config.CollectionUsers).
		Doc(userID).
		Collection(config.CollectionTokens).
		Doc(docID).
		Delete(ctx)
	if err != nil {
		log.Errorm(ctx, "r.client.Delete", err)
		return err
	}
	return nil
}

// NewTokenFirestore ... リポジトリを作成する
func NewTokenFirestore(client *firestore.Client) Token {
	return &tokenFirestore{
		client: client,
	}
}

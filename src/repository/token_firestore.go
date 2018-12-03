package repository

import (
	"context"

	"github.com/aikizoku/push/src/config"
	"github.com/aikizoku/push/src/lib/cloudfirestore"
	"github.com/aikizoku/push/src/lib/log"
	"github.com/aikizoku/push/src/lib/util"
	"github.com/aikizoku/push/src/model"
	"google.golang.org/api/iterator"
)

type tokenFirestore struct {
}

func (r *tokenFirestore) GetMultiToUserID(ctx context.Context, userID string) ([]string, error) {
	var tokens []string
	cli, err := cloudfirestore.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "cloudfirestore.NewClient error: %s", err.Error())
		return tokens, err
	}
	iter := cli.Collection(config.CollectionUsers).Doc(userID).Collection(config.CollectionTokens).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Errorf(ctx, "cli.Get error: %s", err.Error())
			return tokens, err
		}
		var token model.TokenFirestore
		err = doc.DataTo(&token)
		if err != nil {
			log.Errorf(ctx, "dsnp.DataTo error: %s", err.Error())
			return nil, err
		}
		tokens = append(tokens, token.Token)
	}
	return tokens, nil
}

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
		log.Errorf(ctx, "cloudfirestore.NewClient error: %s", err.Error())
		return err
	}
	ret, err := cli.Collection(config.CollectionUsers).Doc(userID).Collection(config.CollectionTokens).Doc(docID).Set(ctx, src)
	if err != nil {
		log.Errorf(ctx, "cli.Set error: %s", err.Error())
		return err
	}
	log.Debugf(ctx, "UpdateTime: %s", ret.UpdateTime)
	return nil
}

// NewTokenFirestore ... TokenFirestoreを作成する
func NewTokenFirestore() Token {
	return &tokenFirestore{}
}

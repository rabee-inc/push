package repository

import (
	"context"

	"github.com/aikizoku/push/src/config"
	"github.com/aikizoku/push/src/lib/cloudfirestore"
	"github.com/aikizoku/push/src/lib/util"
	"github.com/aikizoku/push/src/model"
	"google.golang.org/appengine/log"
)

type tokenFirestore struct {
}

func (r *tokenFirestore) Put(ctx context.Context, userID string, platform string, deviceID string, token string) error {
	docID := model.GenerateTokenDocID(platform, deviceID)
	src := &model.Token{
		Platform:  platform,
		DeviceID:  deviceID,
		Token:     token,
		CreatedAt: util.TimeNow().Unix(),
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

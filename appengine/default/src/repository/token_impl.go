package repository

import (
	"context"

	"cloud.google.com/go/firestore"

	"github.com/rabee-inc/push/appengine/default/src/lib/cloudfirestore"
	"github.com/rabee-inc/push/appengine/default/src/lib/log"
	"github.com/rabee-inc/push/appengine/default/src/lib/util"
	"github.com/rabee-inc/push/appengine/default/src/model"
)

type token struct {
	fCli *firestore.Client
}

func (r *token) Get(ctx context.Context, appID string, userID string, platform string, deviceID string) (string, error) {
	docID := model.GenerateTokenDocID(platform, deviceID)
	docRef := model.TokenRef(r.fCli, appID, userID).Doc(docID)
	dst := &model.Token{}
	_, err := cloudfirestore.Get(ctx, docRef, dst)
	if err != nil {
		log.Errorm(ctx, "cloudfirestore.Get", err)
		return "", err
	}
	return dst.Token, nil
}

func (r *token) ListByUser(ctx context.Context, appID string, userID string) ([]string, error) {
	q := model.TokenRef(r.fCli, appID, userID).Query
	tokens := []*model.Token{}
	err := cloudfirestore.ListByQuery(ctx, q, &tokens)
	if err != nil {
		log.Errorm(ctx, "cloudfirestore.ListByQuery", err)
		return nil, err
	}
	dsts := []string{}
	for _, token := range tokens {
		dsts = append(dsts, token.Token)
	}
	return dsts, nil
}

func (r *token) Put(ctx context.Context, appID string, userID string, platform string, deviceID string, token string) error {
	src := &model.Token{
		Platform:  platform,
		DeviceID:  deviceID,
		Token:     token,
		CreatedAt: util.TimeNowUnix(),
	}
	docID := model.GenerateTokenDocID(platform, deviceID)
	docRef := model.TokenRef(r.fCli, appID, userID).Doc(docID)
	err := cloudfirestore.Set(ctx, docRef, src)
	if err != nil {
		log.Errorm(ctx, "cloudfirestore.Set", err)
		return err
	}
	return nil
}

func (r *token) Delete(ctx context.Context, appID string, userID string, platform string, deviceID string) error {
	docID := model.GenerateTokenDocID(platform, deviceID)
	docRef := model.TokenRef(r.fCli, appID, userID).Doc(docID)
	err := cloudfirestore.Delete(ctx, docRef)
	if err != nil {
		log.Errorm(ctx, "cloudfirestore.Delete", err)
		return err
	}
	return nil
}

// NewToken ... リポジトリを作成する
func NewToken(fCli *firestore.Client) Token {
	return &token{
		fCli: fCli,
	}
}

package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/rabee-inc/go-pkg/cloudfirestore"
	"github.com/rabee-inc/go-pkg/log"
	"github.com/rabee-inc/go-pkg/timeutil"
	"github.com/rabee-inc/push/appengine/push/src/model"
)

type token struct {
	fCli *firestore.Client
}

func (r *token) Get(ctx context.Context, appID string, userID string, platform string, deviceID string) (string, error) {
	docID := model.GenerateTokenDocID(platform, deviceID)
	docRef := model.TokenRef(r.fCli, appID, userID).Doc(docID)
	dst := &model.Token{}
	exist, err := cloudfirestore.Get(ctx, docRef, dst)
	if err != nil {
		log.Error(ctx, err)
		return "", err
	}
	if !exist {
		return "", nil
	}
	return dst.Token, nil
}

func (r *token) ListByUser(ctx context.Context, appID string, userID string) ([]string, error) {
	q := model.TokenRef(r.fCli, appID, userID).Query
	tokens := []*model.Token{}
	err := cloudfirestore.ListByQuery(ctx, q, &tokens)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	dsts := []string{}
	for _, token := range tokens {
		if token.Token == "" {
			continue
		}
		dsts = append(dsts, token.Token)
	}
	return dsts, nil
}

func (r *token) ListByAll(ctx context.Context, appID string) ([]string, error) {
	dsts := []string{}
	var cursor *firestore.DocumentSnapshot
	for {
		tokens := []*model.Token{}
		var dsnp *firestore.DocumentSnapshot
		var err error
		q := r.fCli.CollectionGroup("tokens")
		if cursor != nil {
			q.StartAfter(cursor)
		}
		it := q.Documents(ctx)
		for {
			dsnp, err = it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Error(ctx, err)
				return nil, err
			}
			token := &model.Token{}
			err = dsnp.DataTo(token)
			if err != nil {
				log.Error(ctx, err)
				return nil, err
			}
			tokens = append(tokens, token)
		}
		var nCursor *firestore.DocumentSnapshot
		if len(tokens) == 300 {
			nCursor = dsnp
		}
		for _, token := range tokens {
			dsts = append(dsts, token.Token)
		}
		if nCursor == nil {
			break
		}
		cursor = nCursor
	}
	return dsts, nil
}

func (r *token) Put(ctx context.Context, appID string, userID string, platform string, deviceID string, token string) error {
	src := &model.Token{
		Platform:  platform,
		DeviceID:  deviceID,
		Token:     token,
		CreatedAt: timeutil.NowUnix(),
	}
	docID := model.GenerateTokenDocID(platform, deviceID)
	docRef := model.TokenRef(r.fCli, appID, userID).Doc(docID)
	err := cloudfirestore.Set(ctx, docRef, src)
	if err != nil {
		log.Error(ctx, err)
		return err
	}
	return nil
}

func (r *token) Delete(ctx context.Context, appID string, userID string, platform string, deviceID string) error {
	docID := model.GenerateTokenDocID(platform, deviceID)
	docRef := model.TokenRef(r.fCli, appID, userID).Doc(docID)
	err := cloudfirestore.Delete(ctx, docRef)
	if err != nil {
		log.Error(ctx, err)
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

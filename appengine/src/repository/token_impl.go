package repository

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/rabee-inc/go-pkg/cloudfirestore"
	"github.com/rabee-inc/go-pkg/log"
	"github.com/rabee-inc/push/appengine/src/model"
	"google.golang.org/api/iterator"
)

type token struct {
	cFirestore *firestore.Client
}

func NewToken(
	cFirestore *firestore.Client,
) Token {
	return &token{
		cFirestore,
	}
}

func (r *token) Get(
	ctx context.Context,
	appID string,
	userID string,
	tokenID string,
) (*model.Token, error) {
	docRef := model.TokenRef(r.cFirestore, appID, userID).Doc(tokenID)
	dst := &model.Token{}
	exist, err := cloudfirestore.Get(ctx, docRef, dst)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	if !exist {
		return nil, nil
	}
	return dst, nil
}

func (r *token) List(
	ctx context.Context,
	appID string,
	userID string,
) ([]*model.Token, error) {
	q := model.TokenRef(r.cFirestore, appID, userID).Query
	dsts := []*model.Token{}
	err := cloudfirestore.ListByQuery(ctx, q, &dsts)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	return dsts, nil
}

func (r *token) ListAll(
	ctx context.Context,
	appID string,
) ([]*model.Token, error) {
	dsts := []*model.Token{}
	var cursor *firestore.DocumentSnapshot
	for {
		tokens := []*model.Token{}
		var dsnp *firestore.DocumentSnapshot
		var err error
		q := r.cFirestore.CollectionGroup("tokens")
		if cursor != nil {
			q.StartAfter(cursor)
		}
		it := q.Documents(ctx)
		for {
			dsnp, err = it.Next()
			if errors.Is(err, iterator.Done) {
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
		if nCursor == nil {
			break
		}
		cursor = nCursor
	}
	return dsts, nil
}

func (r *token) Set(
	ctx context.Context,
	appID string,
	userID string,
	src *model.Token,
) error {
	docRef := model.TokenRef(r.cFirestore, appID, userID).Doc(src.ID)
	err := cloudfirestore.Set(ctx, docRef, src)
	if err != nil {
		log.Error(ctx, err)
		return err
	}
	return nil
}

func (r *token) Delete(
	ctx context.Context,
	appID string,
	userID string,
	src *model.Token,
) error {
	docRef := model.TokenRef(r.cFirestore, appID, userID).Doc(src.ID)
	err := cloudfirestore.Delete(ctx, docRef)
	if err != nil {
		log.Error(ctx, err)
		return err
	}
	return nil
}

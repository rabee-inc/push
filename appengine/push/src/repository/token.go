package repository

import (
	"context"

	"github.com/rabee-inc/push/appengine/push/src/model"
)

type Token interface {
	Get(
		ctx context.Context,
		appID string,
		userID string,
		tokenID string,
	) (*model.Token, error)

	List(
		ctx context.Context,
		appID string,
		userID string,
	) ([]*model.Token, error)

	ListAll(
		ctx context.Context,
		appID string,
	) ([]*model.Token, error)

	Set(
		ctx context.Context,
		appID string,
		userID string,
		src *model.Token,
	) error

	Delete(
		ctx context.Context,
		appID string,
		userID string,
		src *model.Token,
	) error
}

package repository

import (
	"context"

	"github.com/rabee-inc/push/appengine/push/src/config"
	"github.com/rabee-inc/push/appengine/push/src/model"
)

type Reserve interface {
	Get(
		ctx context.Context,
		appID string,
		reserveID string,
	) (*model.Reserve, error)

	List(
		ctx context.Context,
		appID string,
		query *ReserveListQuery,
	) ([]*model.Reserve, string, error)

	Create(
		ctx context.Context,
		appID string,
		src *model.Reserve,
	) error

	Update(
		ctx context.Context,
		appID string,
		src *model.Reserve,
	) error
}

type ReserveListQuery struct {
	OverdueReserved    bool
	FilterUnManaged    bool
	FilterStatuses     []config.ReserveStatus
	SortReservedAtDesc bool
	Limit              int
	Cursor             string
}

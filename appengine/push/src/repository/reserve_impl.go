package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/rabee-inc/go-pkg/cloudfirestore"
	"github.com/rabee-inc/go-pkg/log"
	"github.com/rabee-inc/go-pkg/timeutil"
	"github.com/rabee-inc/push/appengine/push/src/model"
)

type reserve struct {
	cFirestore *firestore.Client
}

func NewReserve(
	cFirestore *firestore.Client,
) Reserve {
	return &reserve{
		cFirestore,
	}
}

func (r *reserve) Get(
	ctx context.Context,
	appID string,
	reserveID string,
) (*model.Reserve, error) {
	docRef := model.ReserveRef(r.cFirestore, appID).Doc(reserveID)
	dst := &model.Reserve{}
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

func (r *reserve) List(
	ctx context.Context,
	appID string,
	query *ReserveListQuery,
) ([]*model.Reserve, string, error) {
	q := model.ReserveRef(r.cFirestore, appID).Query
	if query.OverdueReserved {
		q = q.Where("reserved_at", "<=", timeutil.NowUnix())
	}
	if query.FilterUnManaged {
		q = q.Where("unmanaged", "==", false)
	}
	if len(query.FilterStatuses) > 0 {
		q = q.Where("status", "in", query.FilterStatuses)
	}
	if query.SortReservedAtDesc {
		q = q.OrderBy("reserved_at", firestore.Desc)
	}
	var dsnp *firestore.DocumentSnapshot
	var err error
	if query.Cursor != "" {
		dsnp, err = model.ReserveRef(r.cFirestore, appID).Doc(query.Cursor).Get(ctx)
		if err != nil {
			log.Error(ctx, err)
			return nil, "", err
		}
	}
	dsts := []*model.Reserve{}
	nDsnp, err := cloudfirestore.ListByQueryCursor(ctx, q, query.Limit, dsnp, &dsts)
	if err != nil {
		log.Error(ctx, err)
		return nil, "", err
	}
	var nCursor string
	if nDsnp != nil {
		nCursor = nDsnp.Ref.ID
	}
	return dsts, nCursor, nil
}

func (r *reserve) Create(
	ctx context.Context,
	appID string,
	src *model.Reserve,
) error {
	colRef := model.ReserveRef(r.cFirestore, appID)
	err := cloudfirestore.Create(ctx, colRef, src)
	if err != nil {
		log.Error(ctx, err)
		return err
	}
	return nil
}

func (r *reserve) Update(
	ctx context.Context,
	appID string,
	src *model.Reserve,
) error {
	src.UpdatedAt = timeutil.NowUnix()
	docRef := model.ReserveRef(r.cFirestore, appID).Doc(src.ID)
	err := cloudfirestore.Set(ctx, docRef, src)
	if err != nil {
		log.Error(ctx, err)
		return err
	}
	return nil
}

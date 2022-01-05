package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/rabee-inc/go-pkg/cloudfirestore"
	"github.com/rabee-inc/go-pkg/log"

	"github.com/rabee-inc/push/appengine/push/src/config"
	"github.com/rabee-inc/push/appengine/push/src/model"
)

type reserve struct {
	fCli *firestore.Client
}

func (r *reserve) Get(
	ctx context.Context,
	appID string,
	reserveID string) (*model.Reserve, error) {
	docRef := model.ReserveRef(r.fCli, appID).Doc(reserveID)
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

func (r *reserve) ListByCursor(
	ctx context.Context,
	appID string,
	limit int,
	cursor string) ([]*model.Reserve, string, error) {
	q := model.ReserveRef(r.fCli, appID).
		Where("unmanaged", "==", false).
		OrderBy("reserved_at", firestore.Desc)
	var dsnp *firestore.DocumentSnapshot
	var err error
	if cursor != "" {
		dsnp, err = model.ReserveRef(r.fCli, appID).Doc(cursor).Get(ctx)
		if err != nil {
			log.Error(ctx, err)
			return nil, "", err
		}
	}
	dsts := []*model.Reserve{}
	nDsnp, err := cloudfirestore.ListByQueryCursor(ctx, q, limit, dsnp, &dsts)
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

func (r *reserve) ListBySend(
	ctx context.Context,
	appID string,
	now int64,
	limit int,
	cursor *firestore.DocumentSnapshot) ([]*model.Reserve, *firestore.DocumentSnapshot, error) {
	q := model.ReserveRef(r.fCli, appID).
		Where("reserved_at", "<=", now).
		Where("status", "in", []config.ReserveStatus{config.ReserveStatusReserved, config.ReserveStatusFailure})
	var err error
	dsts := []*model.Reserve{}
	nCursor, err := cloudfirestore.ListByQueryCursor(ctx, q, limit, cursor, &dsts)
	if err != nil {
		log.Error(ctx, err)
		return nil, nil, err
	}
	return dsts, nCursor, nil
}

func (r *reserve) Create(
	ctx context.Context,
	appID string,
	userIDs []string,
	msg *model.Message,
	reservedAt int64,
	status config.ReserveStatus,
	unmanaged bool,
	createdAt int64) (*model.Reserve, error) {
	src := &model.Reserve{
		UserIDs:    userIDs,
		Message:    msg,
		ReservedAt: reservedAt,
		Status:     status,
		Unmanaged:  unmanaged,
		CreatedAt:  createdAt,
		UpdatedAt:  createdAt,
	}
	colRef := model.ReserveRef(r.fCli, appID)
	err := cloudfirestore.Create(ctx, colRef, src)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	return src, nil
}

func (r *reserve) Update(
	ctx context.Context,
	appID string,
	src *model.Reserve,
	updatedAt int64) (*model.Reserve, error) {
	src.UpdatedAt = updatedAt
	docRef := model.ReserveRef(r.fCli, appID).Doc(src.ID)
	err := cloudfirestore.Set(ctx, docRef, src)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	return src, nil
}

// NewReserve ... リポジトリを作成する
func NewReserve(fCli *firestore.Client) Reserve {
	return &reserve{
		fCli: fCli,
	}
}

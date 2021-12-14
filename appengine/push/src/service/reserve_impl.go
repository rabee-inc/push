package service

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"

	"github.com/rabee-inc/go-pkg/log"
	"github.com/rabee-inc/go-pkg/timeutil"
	"github.com/rabee-inc/push/appengine/push/src/config"
	"github.com/rabee-inc/push/appengine/push/src/model"
	"github.com/rabee-inc/push/appengine/push/src/repository"
)

type reserve struct {
	rRepo repository.Reserve
	fCli  *firestore.Client
}

func (s *reserve) Get(
	ctx context.Context,
	appID string,
	reserveID string) (*model.Reserve, error) {
	dst, err := s.rRepo.Get(ctx, appID, reserveID)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	if dst == nil {
		err = log.Warningc(ctx, http.StatusNotFound, "not exist reserve: %s", reserveID)
		return nil, err
	}
	return dst, nil
}

func (s *reserve) List(
	ctx context.Context,
	appID string,
	limit int,
	cursor string) ([]*model.Reserve, string, error) {
	dsts, nCursor, err := s.rRepo.ListByCursor(ctx, appID, limit, cursor)
	if err != nil {
		log.Error(ctx, err)
		return nil, "", err
	}
	return dsts, nCursor, nil
}

func (s *reserve) Create(
	ctx context.Context,
	appID string,
	userIDs []string,
	msg *model.Message,
	reservedAt int64,
	unmanaged bool) (*model.Reserve, error) {
	now := timeutil.NowUnix()
	dst, err := s.rRepo.Create(ctx, appID, userIDs, msg, reservedAt, config.ReserveStatusReserved, unmanaged, now)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	return dst, nil
}

func (s *reserve) Update(
	ctx context.Context,
	appID string,
	reserveID string,
	userIDs []string,
	msg *model.Message,
	reservedAt int64,
	status config.ReserveStatus) (*model.Reserve, error) {
	src, err := s.rRepo.Get(ctx, appID, reserveID)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	if src == nil {
		err = log.Warningc(ctx, http.StatusNotFound, "not exist reserve: %s", reserveID)
		return nil, err
	}
	src.UserIDs = userIDs
	src.Message = msg
	src.ReservedAt = reservedAt
	src.Status = status
	now := timeutil.NowUnix()
	dst, err := s.rRepo.Update(ctx, appID, src, now)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	return dst, nil
}

// NewReserve ... サービスを作成する
func NewReserve(
	rRepo repository.Reserve,
	fCli *firestore.Client) Reserve {
	return &reserve{
		rRepo: rRepo,
		fCli:  fCli,
	}
}

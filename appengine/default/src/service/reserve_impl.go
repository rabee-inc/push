package service

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"

	"github.com/rabee-inc/push/appengine/default/src/config"
	"github.com/rabee-inc/push/appengine/default/src/lib/log"
	"github.com/rabee-inc/push/appengine/default/src/lib/util"
	"github.com/rabee-inc/push/appengine/default/src/model"
	"github.com/rabee-inc/push/appengine/default/src/repository"
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
		log.Errorm(ctx, "s.rRepo.Get", err)
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
		log.Errorm(ctx, "s.rRepo.ListByCursor", err)
		return nil, "", err
	}
	return dsts, nCursor, nil
}

func (s *reserve) Create(
	ctx context.Context,
	appID string,
	msg *model.Message,
	reservedAt int64) (*model.Reserve, error) {
	now := util.TimeNowUnix()
	dst, err := s.rRepo.Create(ctx, appID, msg, reservedAt, config.ReserveStatusReserved, now)
	if err != nil {
		log.Errorm(ctx, "s.rRepo.Create", err)
		return nil, err
	}
	return dst, nil
}

func (s *reserve) Update(
	ctx context.Context,
	appID string,
	reserveID string,
	msg *model.Message,
	reservedAt int64,
	status config.ReserveStatus) (*model.Reserve, error) {
	src, err := s.rRepo.Get(ctx, appID, reserveID)
	if err != nil {
		log.Errorm(ctx, "s.rRepo.Get", err)
		return nil, err
	}
	if src == nil {
		err = log.Warningc(ctx, http.StatusNotFound, "not exist reserve: %s", reserveID)
		return nil, err
	}
	src.ReservedAt = reservedAt
	src.Message = msg
	src.Status = status
	now := util.TimeNowUnix()
	dst, err := s.rRepo.Update(ctx, appID, src, now)
	if err != nil {
		log.Errorm(ctx, "s.rRepo.Update", err)
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

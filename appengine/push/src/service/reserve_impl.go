package service

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/rabee-inc/go-pkg/log"
	"github.com/rabee-inc/go-pkg/util"
	"github.com/rabee-inc/push/appengine/push/src/model"
	"github.com/rabee-inc/push/appengine/push/src/repository"
)

type reserve struct {
	rReserve   repository.Reserve
	cFirestore *firestore.Client
}

func NewReserve(
	rReserve repository.Reserve,
	cFirestore *firestore.Client,
) Reserve {
	return &reserve{
		rReserve,
		cFirestore,
	}
}

func (s *reserve) Get(
	ctx context.Context,
	param *ReserveGetInput,
) (*ReserveGetOutput, error) {
	reserve, err := s.rReserve.Get(ctx, param.AppID, param.ReserveID)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	if reserve == nil {
		err = log.Warningc(ctx, http.StatusNotFound, "not found reserve: %s", param.ReserveID)
		return nil, err
	}
	return &ReserveGetOutput{
		Reserve: reserve,
	}, nil
}

func (s *reserve) List(
	ctx context.Context,
	param *ReserveListInput,
) (*ReserveListOutput, error) {
	reserves, nCursor, err := s.rReserve.List(
		ctx,
		param.AppID,
		&repository.ReserveListQuery{
			FilterUnManaged:    true,
			SortReservedAtDesc: true,
			Limit:              param.Limit,
			Cursor:             param.Cursor,
		},
	)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	return &ReserveListOutput{
		Reserves:   reserves,
		NextCursor: nCursor,
	}, nil
}

func (s *reserve) Create(
	ctx context.Context,
	param *ReserveCreateInput,
) (*ReserveCreateOutput, error) {
	reserve := model.NewReserve(
		param.UserIDs,
		param.Message,
		param.ReservedAt,
		false,
	)
	err := s.rReserve.Create(
		ctx,
		param.AppID,
		reserve,
	)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	return &ReserveCreateOutput{
		Reserve: reserve,
	}, nil
}

func (s *reserve) Update(
	ctx context.Context,
	param *ReserveUpdateInput,
) (*ReserveUpdateOutput, error) {
	reserve, err := s.rReserve.Get(ctx, param.AppID, param.ReserveID)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	if reserve == nil {
		err = log.Warningc(ctx, http.StatusNotFound, "not found reserve: %s", param.ReserveID)
		return nil, err
	}

	if param.UserIDs != nil {
		util.AssignIfNotNil(&reserve.UserIDs, &param.UserIDs)
	}
	util.AssignIfNotNil(reserve.Message, param.Message)
	util.AssignIfNotNil(&reserve.ReservedAt, param.ReservedAt)
	util.AssignIfNotNil(&reserve.Status, param.Status)
	err = s.rReserve.Update(
		ctx,
		param.AppID,
		reserve,
	)
	if err != nil {
		log.Error(ctx, err)
		return nil, err
	}
	return &ReserveUpdateOutput{
		Reserve: reserve,
	}, nil
}

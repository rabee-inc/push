package service

import (
	"context"

	"github.com/rabee-inc/push/appengine/push/src/config"
	"github.com/rabee-inc/push/appengine/push/src/model"
)

type Reserve interface {
	Get(
		ctx context.Context,
		param *ReserveGetInput,
	) (*ReserveGetOutput, error)

	List(
		ctx context.Context,
		param *ReserveListInput,
	) (*ReserveListOutput, error)

	Create(
		ctx context.Context,
		param *ReserveCreateInput,
	) (*ReserveCreateOutput, error)

	Update(
		ctx context.Context,
		param *ReserveUpdateInput,
	) (*ReserveUpdateOutput, error)
}

type ReserveGetInput struct {
	AppID     string `json:"app_id"     validate:"required"`
	ReserveID string `json:"reserve_id" validate:"required"`
}

type ReserveGetOutput struct {
	Reserve *model.Reserve `json:"reserve"`
}

type ReserveListInput struct {
	AppID  string `json:"app_id" validate:"required"`
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type ReserveListOutput struct {
	Reserves   []*model.Reserve `json:"reserves"`
	NextCursor string           `json:"next_cursor"`
}

type ReserveCreateInput struct {
	AppID      string         `json:"app_id"      validate:"required"`
	UserIDs    []string       `json:"user_ids"`
	Message    *model.Message `json:"message"     validate:"required"`
	ReservedAt int64          `json:"reserved_at" validate:"required"`
	UnManaged  bool           `json:"unmanaged"`
}

type ReserveCreateOutput struct {
	Reserve *model.Reserve `json:"reserve"`
}

type ReserveUpdateInput struct {
	AppID      string                `json:"app_id"      validate:"required"`
	ReserveID  string                `json:"reserve_id"  validate:"required"`
	UserIDs    []string              `json:"user_ids"`
	Message    *model.Message        `json:"message"     validate:"required"`
	ReservedAt *int64                `json:"reserved_at" validate:"required"`
	Status     *config.ReserveStatus `json:"status"      validate:"required"`
}

type ReserveUpdateOutput struct {
	Reserve *model.Reserve `json:"reserve"`
}

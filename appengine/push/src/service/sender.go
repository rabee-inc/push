package service

import (
	"context"

	"github.com/rabee-inc/push/appengine/push/src/model"
	"github.com/rabee-inc/push/appengine/push/src/model/input"
)

type Sender interface {
	AllUsers(
		ctx context.Context,
		param *SenderAllUsersInput,
	) (*SenderAllUsersOutput, error)

	Users(
		ctx context.Context,
		param *input.WorkerSendUsers,
	) error

	User(
		ctx context.Context,
		param *input.WorkerSendUser,
	) error

	Reserved(
		ctx context.Context,
		param *SenderReservedInput,
	) error
}

type SenderAllUsersInput struct {
	AppID   string         `json:"app_id"  validate:"required"`
	PushID  string         `json:"push_id" validate:"required"`
	Message *model.Message `json:"message" validate:"required"`
}

type SenderAllUsersOutput struct {
	Success bool `json:"success"`
}

type SenderReservedInput struct {
	AppID string `form:"app_id" validate:"required"`
}

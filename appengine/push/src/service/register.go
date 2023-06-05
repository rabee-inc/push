package service

import (
	"context"
)

type Register interface {
	Entry(
		ctx context.Context,
		param *RegisterEntryInput,
	) (*RegisterEntryOutput, error)

	Leave(
		ctx context.Context,
		param *RegisterLeaveInput,
	) (*RegisterLeaveOutput, error)
}

type RegisterEntryInput struct {
	AppID    string `json:"app_id"    validate:"required"`
	UserID   string `json:"user_id"   validate:"required"`
	Platform string `json:"platform"  validate:"required"`
	DeviceID string `json:"device_id"`
	Token    string `json:"token"     validate:"required"`
}

type RegisterEntryOutput struct {
	Success bool `json:"success"`
}

type RegisterLeaveInput struct {
	AppID    string `json:"app_id"    validate:"required"`
	UserID   string `json:"user_id"   validate:"required"`
	Platform string `json:"platform"  validate:"required"`
	DeviceID string `json:"device_id"`
}

type RegisterLeaveOutput struct {
	Success bool `json:"success"`
}

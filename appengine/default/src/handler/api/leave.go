package api

import (
	"context"
	"encoding/json"

	"gopkg.in/go-playground/validator.v9"

	"github.com/rabee-inc/push/appengine/default/src/service"
)

// LeaveAction ... 離脱のアクション
type LeaveAction struct {
	rSvc service.Register
}

type leaveParams struct {
	AppID    string `json:"app_id"    validate:"required"`
	UserID   string `json:"user_id"   validate:"required"`
	Platform string `json:"platform"  validate:"required"`
	DeviceID string `json:"device_id"`
}

type leaveResponse struct {
	Success bool `json:"success"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *LeaveAction) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params leaveParams
	err := json.Unmarshal(*msg, &params)
	if err != nil {
		return params, err
	}

	// Validation
	v := validator.New()
	if err := v.Struct(params); err != nil {
		return params, err
	}
	return params, nil
}

// Exec ... 処理をする
func (h *LeaveAction) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	ps := params.(leaveParams)

	err := h.rSvc.DeleteToken(ctx, ps.AppID, ps.UserID, ps.Platform, ps.DeviceID)
	if err != nil {
		return nil, err
	}

	return leaveResponse{
		Success: true,
	}, nil
}

// NewLeaveAction ... アクションを作成する
func NewLeaveAction(rSvc service.Register) *LeaveAction {
	return &LeaveAction{
		rSvc: rSvc,
	}
}

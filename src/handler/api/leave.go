package api

import (
	"context"
	"encoding/json"

	"github.com/rabee-inc/push/src/service"
	"gopkg.in/go-playground/validator.v9"
)

// LeaveAction ... 離脱のアクション
type LeaveAction struct {
	Svc service.Register
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
	// パラメータを取得
	ps := params.(leaveParams)

	err := h.Svc.DeleteToken(ctx, ps.AppID, ps.UserID, ps.Platform, ps.DeviceID)
	if err != nil {
		return nil, err
	}

	return leaveResponse{
		Success: true,
	}, nil
}

// NewLeaveAction ... アクションを作成する
func NewLeaveAction(svc service.Register) *LeaveAction {
	return &LeaveAction{
		Svc: svc,
	}
}

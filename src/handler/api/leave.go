package api

import (
	"context"
	"encoding/json"

	"github.com/rabee-inc/push/src/lib/log"
	"github.com/rabee-inc/push/src/service"
	"gopkg.in/go-playground/validator.v9"
)

// LeaveHandler ... 離脱のハンドラ
type LeaveHandler struct {
	Svc service.Register
}

type leaveParams struct {
	UserID   string `json:"user_id"   validate:"required"`
	Platform string `json:"platform"  validate:"required"`
	DeviceID string `json:"device_id"`
}

type leaveResponse struct {
	Success bool `json:"success"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *LeaveHandler) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
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
func (h *LeaveHandler) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	// パラメータを取得
	ps := params.(leaveParams)

	err := h.Svc.DeleteToken(ctx, ps.UserID, ps.Platform, ps.DeviceID)
	if err != nil {
		log.Errorm(ctx, "h.Svc.DeleteToken", err)
		return nil, err
	}

	return leaveResponse{
		Success: true,
	}, nil
}

// NewLeaveHandler ... ハンドラを作成する
func NewLeaveHandler(svc service.Register) *LeaveHandler {
	return &LeaveHandler{
		Svc: svc,
	}
}

package api

import (
	"context"
	"encoding/json"

	"gopkg.in/go-playground/validator.v9"

	"github.com/rabee-inc/push/appengine/push/src/service"
)

// EntryAction ... エントリーのアクション
type EntryAction struct {
	rSvc service.Register
	v    *validator.Validate
}

type entryParams struct {
	AppID    string `json:"app_id"    validate:"required"`
	UserID   string `json:"user_id"   validate:"required"`
	Platform string `json:"platform"  validate:"required"`
	DeviceID string `json:"device_id"`
	Token    string `json:"token"     validate:"required"`
}

type entryResponse struct {
	Success bool `json:"success"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *EntryAction) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params entryParams
	err := json.Unmarshal(*msg, &params)
	if err != nil {
		return params, err
	}

	// Validation
	if err := h.v.Struct(params); err != nil {
		return params, err
	}
	return params, nil
}

// Exec ... 処理をする
func (h *EntryAction) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	ps := params.(entryParams)

	err := h.rSvc.SetToken(ctx, ps.AppID, ps.UserID, ps.Platform, ps.DeviceID, ps.Token)
	if err != nil {
		return nil, err
	}

	return entryResponse{
		Success: true,
	}, nil
}

// NewEntryAction ... アクションを作成する
func NewEntryAction(rSvc service.Register) *EntryAction {
	v := validator.New()
	return &EntryAction{
		rSvc: rSvc,
		v:    v,
	}
}

package api

import (
	"context"
	"encoding/json"

	"github.com/rabee-inc/push/src/lib/log"
	"github.com/rabee-inc/push/src/service"
	"github.com/go-playground/validator"
)

// EntryHandler ... エントリーのハンドラ
type EntryHandler struct {
	Svc service.Entry
}

type entryParams struct {
	UserID   string `json:"user_id" validate:"required"`
	Platform string `json:"platform" validate:"required"`
	DeviceID string `json:"device_id"`
	Token    string `json:"token" validate:"required"`
}

type entryResponse struct {
	Success bool `json:"success"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *EntryHandler) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params entryParams
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
func (h *EntryHandler) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	// パラメータを取得
	ps := params.(entryParams)

	err := h.Svc.Token(ctx, ps.UserID, ps.Platform, ps.DeviceID, ps.Token)
	if err != nil {
		log.Errorm(ctx, "h.Svc.Token", err)
		return nil, err
	}

	return entryResponse{
		Success: true,
	}, nil
}

// NewEntryHandler ... EntryHandlerを作成する
func NewEntryHandler(svc service.Entry) *EntryHandler {
	return &EntryHandler{
		Svc: svc,
	}
}

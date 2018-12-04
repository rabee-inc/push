package api

import (
	"context"
	"encoding/json"

	"github.com/aikizoku/push/src/lib/log"
	"github.com/aikizoku/push/src/service"
)

// EntryHandler ... エントリーのハンドラ
type EntryHandler struct {
	Svc service.Entry
}

type entryParams struct {
	UserID   string `json:"user_id"`
	Platform string `json:"platform"`
	DeviceID string `json:"device_id"`
	Token    string `json:"token"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *EntryHandler) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params entryParams
	err := json.Unmarshal(*msg, &params)
	return params, err
}

// Exec ... 処理をする
func (h *EntryHandler) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	// パラメータを取得
	ps := params.(entryParams)

	err := h.Svc.Token(ctx, ps.UserID, ps.Platform, ps.DeviceID, ps.Token)
	if err != nil {
		log.Errorf(ctx, "h.Svc.Token error: %s", err.Error())
		return nil, err
	}

	return struct {
		Success bool `json:"success"`
	}{
		Success: true,
	}, nil
}

// NewEntryHandler ... EntryHandlerを作成する
func NewEntryHandler(svc service.Entry) *EntryHandler {
	return &EntryHandler{
		Svc: svc,
	}
}

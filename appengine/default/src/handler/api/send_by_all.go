package api

import (
	"context"
	"encoding/json"

	"gopkg.in/go-playground/validator.v9"

	"github.com/rabee-inc/push/appengine/default/src/model"
	"github.com/rabee-inc/push/appengine/default/src/service"
)

// SendByAllAction ...  送信のアクション
type SendByAllAction struct {
	sSvc service.Sender
}

type sendByAllParams struct {
	AppID   string         `json:"app_id"   validate:"required"`
	Message *model.Message `json:"message"  validate:"required"`
}

type sendByAllResponse struct {
	Success bool `json:"success"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *SendByAllAction) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params sendByAllParams
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
func (h *SendByAllAction) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	ps := params.(sendByAllParams)

	err := h.sSvc.MessageByAll(ctx, ps.AppID, ps.Message)
	if err != nil {
		return nil, err
	}

	return sendByAllResponse{
		Success: true,
	}, nil
}

// NewSendByAllAction ... SendByAllActionを作成する
func NewSendByAllAction(sSvc service.Sender) *SendByAllAction {
	return &SendByAllAction{
		sSvc: sSvc,
	}
}

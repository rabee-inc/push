package api

import (
	"context"
	"encoding/json"

	"gopkg.in/go-playground/validator.v9"

	"github.com/rabee-inc/push/appengine/push/src/model"
	"github.com/rabee-inc/push/appengine/push/src/service"
)

// SendByAllUsersAction ...  全員送信のアクション
type SendByAllUsersAction struct {
	sSvc service.Sender
	v    *validator.Validate
}

type sendByAllUsersParams struct {
	AppID   string         `json:"app_id"  validate:"required"`
	PushID  string         `json:"push_id" validate:"required"`
	Message *model.Message `json:"message" validate:"required"`
}

type sendByAllUsersResponse struct {
	Success bool `json:"success"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *SendByAllUsersAction) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params sendByAllUsersParams
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
func (h *SendByAllUsersAction) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	ps := params.(sendByAllUsersParams)

	err := h.sSvc.AllUsers(ctx, ps.AppID, ps.PushID, ps.Message)
	if err != nil {
		return nil, err
	}

	return sendByAllUsersResponse{
		Success: true,
	}, nil
}

// NewSendByAllUsersAction ... SendByAllUsersActionを作成する
func NewSendByAllUsersAction(sSvc service.Sender) *SendByAllUsersAction {
	v := validator.New()
	return &SendByAllUsersAction{
		sSvc: sSvc,
		v:    v,
	}
}

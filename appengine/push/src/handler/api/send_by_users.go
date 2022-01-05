package api

import (
	"context"
	"encoding/json"

	"github.com/rabee-inc/go-pkg/cloudtasks"
	"gopkg.in/go-playground/validator.v9"

	"github.com/rabee-inc/push/appengine/push/src/config"
	"github.com/rabee-inc/push/appengine/push/src/model"
)

// SendByUsersAction ...  送信のアクション
type SendByUsersAction struct {
	tCli *cloudtasks.Client
	v    *validator.Validate
}

type sendByUsersParams struct {
	AppID   string         `json:"app_id"   validate:"required"`
	UserIDs []string       `json:"user_ids" validate:"required"`
	PushID  string         `json:"push_id"  validate:"required"`
	Message *model.Message `json:"message"  validate:"required"`
}

type sendByUsersResponse struct {
	Success bool `json:"success"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *SendByUsersAction) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params sendByUsersParams
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
func (h *SendByUsersAction) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	ps := params.(sendByUsersParams)

	src := &model.CloudTasksParamSendUsers{
		AppID:   ps.AppID,
		UserIDs: ps.UserIDs,
		PushID:  ps.PushID,
		Message: ps.Message,
	}

	err := h.tCli.AddTask(ctx, config.QueueSendUser, "/worker/send/users", src)
	if err != nil {
		return nil, err
	}

	return sendByUsersResponse{
		Success: true,
	}, nil
}

// NewSendByUsersAction ... SendByUsersActionを作成する
func NewSendByUsersAction(tCli *cloudtasks.Client) *SendByUsersAction {
	v := validator.New()
	return &SendByUsersAction{
		tCli: tCli,
		v:    v,
	}
}

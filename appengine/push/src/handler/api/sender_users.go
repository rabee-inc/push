package api

import (
	"context"
	"encoding/json"

	"github.com/rabee-inc/go-pkg/cloudtasks"
	"github.com/rabee-inc/push/appengine/push/src/config"
	"github.com/rabee-inc/push/appengine/push/src/handler"
	"github.com/rabee-inc/push/appengine/push/src/model/input"
)

type SenderUsers struct {
	cTasks *cloudtasks.Client
}

func NewSenderUsers(
	cTasks *cloudtasks.Client,
) *SenderUsers {
	return &SenderUsers{
		cTasks,
	}
}

func (h *SenderUsers) DecodeParams(ctx context.Context, msg *json.RawMessage) (any, error) {
	var param input.WorkerSendUsers
	err := json.Unmarshal(*msg, &param)
	if err != nil {
		return nil, err
	}
	if err = handler.Validate(param); err != nil {
		return nil, err
	}
	return param, nil
}

func (h *SenderUsers) Exec(ctx context.Context, method string, params any) (any, error) {
	param := params.(input.WorkerSendUsers)
	err := h.cTasks.AddTask(ctx, config.QueueSendUser, "/worker/send/users", param)
	if err != nil {
		return nil, err
	}
	return &struct {
		Success bool `json:"success"`
	}{
		Success: true,
	}, nil
}

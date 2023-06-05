package api

import (
	"context"
	"encoding/json"

	"github.com/rabee-inc/push/appengine/push/src/handler"
	"github.com/rabee-inc/push/appengine/push/src/service"
)

type SenderAllUsers struct {
	sSender service.Sender
}

func NewSenderAllUsers(
	sSender service.Sender,
) *SenderAllUsers {
	return &SenderAllUsers{
		sSender,
	}
}

func (h *SenderAllUsers) DecodeParams(ctx context.Context, msg *json.RawMessage) (any, error) {
	var param service.SenderAllUsersInput
	err := json.Unmarshal(*msg, &param)
	if err != nil {
		return nil, err
	}
	if err = handler.Validate(param); err != nil {
		return nil, err
	}
	return param, nil
}

func (h *SenderAllUsers) Exec(ctx context.Context, method string, params any) (any, error) {
	param := params.(service.SenderAllUsersInput)
	output, err := h.sSender.AllUsers(ctx, &param)
	if err != nil {
		return nil, err
	}
	return output, nil
}

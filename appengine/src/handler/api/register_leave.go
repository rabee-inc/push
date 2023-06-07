package api

import (
	"context"
	"encoding/json"

	"github.com/rabee-inc/push/appengine/src/handler"
	"github.com/rabee-inc/push/appengine/src/service"
)

type RegisterLeave struct {
	sRegister service.Register
}

func NewRegisterLeave(
	sRegister service.Register,
) *RegisterLeave {
	return &RegisterLeave{
		sRegister,
	}
}

func (h *RegisterLeave) DecodeParams(ctx context.Context, msg *json.RawMessage) (any, error) {
	var param service.RegisterLeaveInput
	err := json.Unmarshal(*msg, &param)
	if err != nil {
		return nil, err
	}
	if err = handler.Validate(param); err != nil {
		return nil, err
	}
	return param, nil
}

func (h *RegisterLeave) Exec(ctx context.Context, method string, params any) (any, error) {
	param := params.(service.RegisterLeaveInput)
	output, err := h.sRegister.Leave(ctx, &param)
	if err != nil {
		return nil, err
	}
	return output, nil
}

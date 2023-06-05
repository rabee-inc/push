package api

import (
	"context"
	"encoding/json"

	"github.com/rabee-inc/push/appengine/push/src/handler"
	"github.com/rabee-inc/push/appengine/push/src/service"
)

type ReserveGet struct {
	sReserve service.Reserve
}

func NewReserveGet(
	sReserve service.Reserve,
) *ReserveGet {
	return &ReserveGet{
		sReserve,
	}
}

func (h *ReserveGet) DecodeParams(ctx context.Context, msg *json.RawMessage) (any, error) {
	var param service.ReserveGetInput
	err := json.Unmarshal(*msg, &param)
	if err != nil {
		return nil, err
	}
	if err = handler.Validate(param); err != nil {
		return nil, err
	}
	return param, nil
}

func (h *ReserveGet) Exec(ctx context.Context, method string, params any) (any, error) {
	param := params.(service.ReserveGetInput)
	output, err := h.sReserve.Get(ctx, &param)
	if err != nil {
		return nil, err
	}
	return output, nil
}

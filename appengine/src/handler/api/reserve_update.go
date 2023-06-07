package api

import (
	"context"
	"encoding/json"

	"github.com/rabee-inc/push/appengine/src/handler"
	"github.com/rabee-inc/push/appengine/src/service"
)

type ReserveUpdate struct {
	sReserve service.Reserve
}

func NewReserveUpdate(
	sReserve service.Reserve,
) *ReserveUpdate {
	return &ReserveUpdate{
		sReserve,
	}
}

func (h *ReserveUpdate) DecodeParams(ctx context.Context, msg *json.RawMessage) (any, error) {
	var param service.ReserveUpdateInput
	err := json.Unmarshal(*msg, &param)
	if err != nil {
		return nil, err
	}
	if err = handler.Validate(param); err != nil {
		return nil, err
	}
	return param, nil
}

func (h *ReserveUpdate) Exec(ctx context.Context, method string, params any) (any, error) {
	param := params.(service.ReserveUpdateInput)
	output, err := h.sReserve.Update(ctx, &param)
	if err != nil {
		return nil, err
	}
	return output, nil
}

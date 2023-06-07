package api

import (
	"context"
	"encoding/json"

	"github.com/rabee-inc/push/appengine/src/handler"
	"github.com/rabee-inc/push/appengine/src/service"
)

type ReserveCreate struct {
	sReserve service.Reserve
}

func NewReserveCreate(
	sReserve service.Reserve,
) *ReserveCreate {
	return &ReserveCreate{
		sReserve,
	}
}

func (h *ReserveCreate) DecodeParams(ctx context.Context, msg *json.RawMessage) (any, error) {
	var param service.ReserveCreateInput
	err := json.Unmarshal(*msg, &param)
	if err != nil {
		return nil, err
	}
	if err = handler.Validate(param); err != nil {
		return nil, err
	}
	return param, nil
}

func (h *ReserveCreate) Exec(ctx context.Context, method string, params any) (any, error) {
	param := params.(service.ReserveCreateInput)
	output, err := h.sReserve.Create(ctx, &param)
	if err != nil {
		return nil, err
	}
	return output, nil
}

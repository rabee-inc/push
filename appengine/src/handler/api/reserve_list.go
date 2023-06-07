package api

import (
	"context"
	"encoding/json"

	"github.com/rabee-inc/push/appengine/src/handler"
	"github.com/rabee-inc/push/appengine/src/service"
)

type ReserveList struct {
	sReserve service.Reserve
}

func NewReserveList(
	sReserve service.Reserve,
) *ReserveList {
	return &ReserveList{
		sReserve,
	}
}

func (h *ReserveList) DecodeParams(ctx context.Context, msg *json.RawMessage) (any, error) {
	var param service.ReserveListInput
	err := json.Unmarshal(*msg, &param)
	if err != nil {
		return nil, err
	}
	if err = handler.Validate(param); err != nil {
		return nil, err
	}
	return param, nil
}

func (h *ReserveList) Exec(ctx context.Context, method string, params any) (any, error) {
	param := params.(service.ReserveListInput)
	output, err := h.sReserve.List(ctx, &param)
	if err != nil {
		return nil, err
	}
	return output, nil
}

package api

import (
	"context"
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/rabee-inc/push/appengine/src/handler"
	"github.com/rabee-inc/push/appengine/src/service"
)

type RegisterEntry struct {
	sRegister service.Register
}

func NewRegisterEntry(
	sRegister service.Register,
) *RegisterEntry {
	return &RegisterEntry{
		sRegister,
	}
}

func (h *RegisterEntry) DecodeParams(ctx context.Context, msg *json.RawMessage) (any, error) {
	var param service.RegisterEntryInput
	err := json.Unmarshal(*msg, &param)
	if err != nil {
		return nil, err
	}
	if err = handler.Validate(param); err != nil {
		return nil, err
	}
	return param, nil
}

func (h *RegisterEntry) Exec(ctx context.Context, method string, params any) (any, error) {
	param := params.(service.RegisterEntryInput)
	pp.Println(param)
	output, err := h.sRegister.Entry(ctx, &param)
	if err != nil {
		return nil, err
	}
	return output, nil
}

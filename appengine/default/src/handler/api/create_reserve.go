package api

import (
	"context"
	"encoding/json"

	"gopkg.in/go-playground/validator.v9"

	"github.com/rabee-inc/push/appengine/default/src/model"
	"github.com/rabee-inc/push/appengine/default/src/service"
)

// CreateReserveAction ... エントリーのアクション
type CreateReserveAction struct {
	rSvc service.Reserve
}

type createReserveParams struct {
	AppID      string         `json:"app_id"      validate:"required"`
	Message    *model.Message `json:"message"     validate:"required"`
	ReservedAt int64          `json:"reserved_at" validate:"required"`
}

type createReserveResponse struct {
	Reserve *model.Reserve `json:"reserve"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *CreateReserveAction) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params createReserveParams
	err := json.Unmarshal(*msg, &params)
	if err != nil {
		return params, err
	}

	// Validation
	v := validator.New()
	if err := v.Struct(params); err != nil {
		return params, err
	}
	return params, nil
}

// Exec ... 処理をする
func (h *CreateReserveAction) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	ps := params.(createReserveParams)

	dst, err := h.rSvc.Create(ctx, ps.AppID, ps.Message, ps.ReservedAt)
	if err != nil {
		return nil, err
	}

	return createReserveResponse{
		Reserve: dst,
	}, nil
}

// NewCreateReserveAction ... アクションを作成する
func NewCreateReserveAction(rSvc service.Reserve) *CreateReserveAction {
	return &CreateReserveAction{
		rSvc: rSvc,
	}
}

package api

import (
	"context"
	"encoding/json"

	"gopkg.in/go-playground/validator.v9"

	"github.com/rabee-inc/push/appengine/push/src/model"
	"github.com/rabee-inc/push/appengine/push/src/service"
)

// GetReserveAction ... 予約取得のアクション
type GetReserveAction struct {
	rSvc service.Reserve
	v    *validator.Validate
}

type getReserveParams struct {
	AppID     string `json:"app_id"     validate:"required"`
	ReserveID string `json:"reserve_id" validate:"required"`
}

type getReserveResponse struct {
	Reserve *model.Reserve `json:"reserve"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *GetReserveAction) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params getReserveParams
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
func (h *GetReserveAction) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	ps := params.(getReserveParams)

	dst, err := h.rSvc.Get(ctx, ps.AppID, ps.ReserveID)
	if err != nil {
		return nil, err
	}

	return getReserveResponse{
		Reserve: dst,
	}, nil
}

// NewGetReserveAction ... アクションを作成する
func NewGetReserveAction(rSvc service.Reserve) *GetReserveAction {
	v := validator.New()
	return &GetReserveAction{
		rSvc: rSvc,
		v:    v,
	}
}

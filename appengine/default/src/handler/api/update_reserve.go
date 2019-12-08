package api

import (
	"context"
	"encoding/json"

	"gopkg.in/go-playground/validator.v9"

	"github.com/rabee-inc/push/appengine/default/src/config"
	"github.com/rabee-inc/push/appengine/default/src/model"
	"github.com/rabee-inc/push/appengine/default/src/service"
)

// UpdateReserveAction ... エントリーのアクション
type UpdateReserveAction struct {
	rSvc service.Reserve
}

type updateReserveParams struct {
	AppID      string               `json:"app_id"      validate:"required"`
	ReserveID  string               `json:"reserve_id"  validate:"required"`
	Message    *model.Message       `json:"message"     validate:"required"`
	ReservedAt int64                `json:"reserved_at" validate:"required"`
	Status     config.ReserveStatus `json:"status"      validate:"oneof=reserved canceled"`
}

type updateReserveResponse struct {
	Reserve *model.Reserve `json:"reserve"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *UpdateReserveAction) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params updateReserveParams
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
func (h *UpdateReserveAction) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	ps := params.(updateReserveParams)

	dst, err := h.rSvc.Update(ctx, ps.AppID, ps.ReserveID, ps.Message, ps.ReservedAt, ps.Status)
	if err != nil {
		return nil, err
	}

	return updateReserveResponse{
		Reserve: dst,
	}, nil
}

// NewUpdateReserveAction ... アクションを作成する
func NewUpdateReserveAction(rSvc service.Reserve) *UpdateReserveAction {
	return &UpdateReserveAction{
		rSvc: rSvc,
	}
}

package api

import (
	"context"
	"encoding/json"

	"gopkg.in/go-playground/validator.v9"

	"github.com/rabee-inc/push/appengine/push/src/config"
	"github.com/rabee-inc/push/appengine/push/src/model"
	"github.com/rabee-inc/push/appengine/push/src/service"
)

// UpdateReserveAction ... 予約編集のアクション
type UpdateReserveAction struct {
	rSvc service.Reserve
	v    *validator.Validate
}

type updateReserveParams struct {
	AppID      string               `json:"app_id"      validate:"required"`
	ReserveID  string               `json:"reserve_id"  validate:"required"`
	UserIDs    []string             `json:"user_ids"`
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
	if err := h.v.Struct(params); err != nil {
		return params, err
	}
	return params, nil
}

// Exec ... 処理をする
func (h *UpdateReserveAction) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	ps := params.(updateReserveParams)

	dst, err := h.rSvc.Update(ctx, ps.AppID, ps.ReserveID, ps.UserIDs, ps.Message, ps.ReservedAt, ps.Status)
	if err != nil {
		return nil, err
	}

	return updateReserveResponse{
		Reserve: dst,
	}, nil
}

// NewUpdateReserveAction ... アクションを作成する
func NewUpdateReserveAction(rSvc service.Reserve) *UpdateReserveAction {
	v := validator.New()
	return &UpdateReserveAction{
		rSvc: rSvc,
		v:    v,
	}
}

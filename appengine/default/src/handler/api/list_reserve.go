package api

import (
	"context"
	"encoding/json"

	"gopkg.in/go-playground/validator.v9"

	"github.com/rabee-inc/push/appengine/default/src/model"
	"github.com/rabee-inc/push/appengine/default/src/service"
)

// ListReserveAction ... エントリーのアクション
type ListReserveAction struct {
	rSvc service.Reserve
}

type listReserveParams struct {
	AppID  string `json:"app_id" validate:"required"`
	Limit  int    `json:"limit"  validate:"required"`
	Cursor string `json:"cursor" validate:"required"`
}

type listReserveResponse struct {
	Reserves   []*model.Reserve `json:"reserves"`
	NextCursor string           `json:"next_cursor"`
}

// DecodeParams ... 受け取ったJSONパラメータをデコードする
func (h *ListReserveAction) DecodeParams(ctx context.Context, msg *json.RawMessage) (interface{}, error) {
	var params listReserveParams
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
func (h *ListReserveAction) Exec(ctx context.Context, method string, params interface{}) (interface{}, error) {
	ps := params.(listReserveParams)

	dsts, nCursor, err := h.rSvc.List(ctx, ps.AppID, ps.Limit, ps.Cursor)
	if err != nil {
		return nil, err
	}

	return listReserveResponse{
		Reserves:   dsts,
		NextCursor: nCursor,
	}, nil
}

// NewListReserveAction ... アクションを作成する
func NewListReserveAction(rSvc service.Reserve) *ListReserveAction {
	return &ListReserveAction{
		rSvc: rSvc,
	}
}

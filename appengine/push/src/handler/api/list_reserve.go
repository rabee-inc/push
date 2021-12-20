package api

import (
	"context"
	"encoding/json"

	"gopkg.in/go-playground/validator.v9"

	"github.com/rabee-inc/push/appengine/push/src/model"
	"github.com/rabee-inc/push/appengine/push/src/service"
)

// ListReserveAction ... 予約リスト取得のアクション
type ListReserveAction struct {
	rSvc service.Reserve
	v    *validator.Validate
}

type listReserveParams struct {
	AppID  string `json:"app_id" validate:"required"`
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
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
	if err := h.v.Struct(params); err != nil {
		return params, err
	}

	if params.Limit == 0 {
		params.Limit = 16
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
	v := validator.New()
	return &ListReserveAction{
		rSvc: rSvc,
		v:    v,
	}
}

package worker

import (
	"net/http"

	"github.com/rabee-inc/go-pkg/errcode"
	"github.com/rabee-inc/go-pkg/parameter"
	"github.com/rabee-inc/go-pkg/renderer"
	"gopkg.in/go-playground/validator.v9"

	"github.com/rabee-inc/push/appengine/push/src/model"
	"github.com/rabee-inc/push/appengine/push/src/service"
)

// SendHandler ... 送信のハンドラ
type SendHandler struct {
	sSvc service.Sender
	rSvc service.Reserve
	v    *validator.Validate
}

// SendByUsers ... UsersからUserに分割してプッシュ通知を送信する
func (h *SendHandler) SendByUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param model.CloudTasksParamSendUsers
	err := parameter.GetJSON(r, &param)
	if err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, err)
		return
	}

	// Validation
	if err := h.v.Struct(param); err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, err)
		return
	}

	err = h.sSvc.Users(ctx, param.AppID, param.UserIDs, param.PushID, param.Message)
	if err != nil {
		renderer.HandleError(ctx, w, err)
		return
	}
	renderer.Success(ctx, w)
}

// SendByUser ... UserからTokenを引いてプッシュ通知を送信する
func (h *SendHandler) SendByUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param model.CloudTasksParamSendUser
	err := parameter.GetJSON(r, &param)
	if err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, err)
		return
	}

	// Validation
	if err := h.v.Struct(param); err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, err)
		return
	}

	err = h.sSvc.User(ctx, param.AppID, param.UserID, param.PushID, param.Message)
	if err != nil {
		renderer.HandleError(ctx, w, err)
		return
	}
	renderer.Success(ctx, w)
}

// SendByReserved ... 予約プッシュ通知を送信する
func (h *SendHandler) SendByReserved(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param struct {
		AppID string `form:"app_id" validate:"required"`
	}
	err := parameter.GetForms(ctx, r, &param)
	if err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, err)
		return
	}

	// Validation
	if err := h.v.Struct(param); err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, err)
		return
	}

	err = h.sSvc.Reserved(ctx, param.AppID)
	if err != nil {
		renderer.HandleError(ctx, w, err)
		return
	}
	renderer.Success(ctx, w)
}

// NewSendHandler ... SendHandlerを作成する
func NewSendHandler(sSvc service.Sender, rSvc service.Reserve) *SendHandler {
	v := validator.New()
	return &SendHandler{
		sSvc: sSvc,
		rSvc: rSvc,
		v:    v,
	}
}

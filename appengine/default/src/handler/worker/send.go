package worker

import (
	"net/http"

	"gopkg.in/go-playground/validator.v9"

	"github.com/rabee-inc/push/appengine/default/src/lib/errcode"
	"github.com/rabee-inc/push/appengine/default/src/lib/parameter"
	"github.com/rabee-inc/push/appengine/default/src/lib/renderer"
	"github.com/rabee-inc/push/appengine/default/src/model"
	"github.com/rabee-inc/push/appengine/default/src/service"
)

// SendHandler ... サンプルのハンドラ定義
type SendHandler struct {
	sSvc service.Sender
	rSvc service.Reserve
}

// SendByUsers ... UsersからUserに分割してプッシュ通知を送信する
func (h *SendHandler) SendByUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param model.CloudTasksParamSendUsers
	err := parameter.GetJSON(r, &param)
	if err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, "parameter.GetJSON", err)
		return
	}

	// Validation
	v := validator.New()
	if err := v.Struct(param); err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, "v.Struct", err)
		return
	}

	err = h.sSvc.Users(ctx, param.AppID, param.UserIDs, param.Message)
	if err != nil {
		renderer.HandleError(ctx, w, "h.sSvc.Users", err)
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
		renderer.HandleError(ctx, w, "parameter.GetJSON", err)
		return
	}

	// Validation
	v := validator.New()
	if err := v.Struct(param); err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, "v.Struct", err)
		return
	}

	err = h.sSvc.User(ctx, param.AppID, param.UserID, param.Message)
	if err != nil {
		renderer.HandleError(ctx, w, "h.sSvc.User", err)
		return
	}
	renderer.Success(ctx, w)
}

// SendByReserved ... 予約プッシュ通知を送信する
func (h *SendHandler) SendByReserved(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param struct {
		AppID string `json:"app_id"`
	}
	err := parameter.GetJSON(r, &param)
	if err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, "parameter.GetJSON", err)
		return
	}

	// Validation
	v := validator.New()
	if err := v.Struct(param); err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, "v.Struct", err)
		return
	}

	err = h.sSvc.Reserved(ctx, param.AppID)
	if err != nil {
		renderer.HandleError(ctx, w, "h.sSvc.Reserved", err)
		return
	}
	renderer.Success(ctx, w)
}

// NewSendHandler ... SendHandlerを作成する
func NewSendHandler(sSvc service.Sender, rSvc service.Reserve) *SendHandler {
	return &SendHandler{
		sSvc: sSvc,
		rSvc: rSvc,
	}
}

package worker

import (
	"net/http"

	"github.com/rabee-inc/push/src/lib/errcode"
	"github.com/rabee-inc/push/src/lib/parameter"
	"github.com/rabee-inc/push/src/lib/renderer"
	"github.com/rabee-inc/push/src/model"
	"github.com/rabee-inc/push/src/service"
	"gopkg.in/go-playground/validator.v9"
)

// SendHandler ... サンプルのハンドラ定義
type SendHandler struct {
	Svc service.Sender
}

// SendUserIDs ... UserIDsからUserIDに分割してプッシュ通知を送信する
func (h *SendHandler) SendUserIDs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param model.TaskQueueParamSendUserIDs
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

	err = h.Svc.MessageByUserIDs(ctx, param.AppID, param.UserIDs, param.Message)
	if err != nil {
		renderer.HandleError(ctx, w, "h.Svc.MessageByUserIDs", err)
		return
	}
	renderer.Success(w)
}

// SendUserID ... UserIDからTokenを引いてプッシュ通知を送信する
func (h *SendHandler) SendUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param model.TaskQueueParamSendUserID
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

	err = h.Svc.MessageByUserID(ctx, param.AppID, param.UserID, param.Message)
	if err != nil {
		renderer.HandleError(ctx, w, "h.Svc.MessageByUserID", err)
		return
	}
	renderer.Success(w)
}

// SendToken ... Tokenでプッシュ通知を送信する
func (h *SendHandler) SendToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param model.TaskQueueParamSendToken
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

	err = h.Svc.MessageByToken(ctx, param.Token, param.Message)
	if err != nil {
		renderer.HandleError(ctx, w, "h.Svc.MessageByToken", err)
		return
	}
	renderer.Success(w)
}

// NewSendHandler ... SendHandlerを作成する
func NewSendHandler(svc service.Sender) *SendHandler {
	return &SendHandler{
		Svc: svc,
	}
}

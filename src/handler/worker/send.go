package worker

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rabee-inc/push/src/handler"
	"github.com/rabee-inc/push/src/lib/log"
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
	err := handler.GetJSON(r, &param)
	if err != nil {
		h.handleError(ctx, w, http.StatusBadRequest, "handler.GetJSON: %s", err.Error())
		return
	}

	// Validation
	v := validator.New()
	if err := v.Struct(param); err != nil {
		h.handleError(ctx, w, http.StatusBadRequest, "v.Struct: %s", err.Error())
		return
	}

	err = h.Svc.MessageByUserIDs(ctx, param.AppID, param.UserIDs, param.Message)
	if err != nil {
		h.handleError(ctx, w, http.StatusBadRequest, "h.Svc.MessageByUserIDs: %s", err.Error())
		return
	}
	handler.RenderSuccess(w)
}

// SendUserID ... UserIDからTokenを引いてプッシュ通知を送信する
func (h *SendHandler) SendUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param model.TaskQueueParamSendUserID
	err := handler.GetJSON(r, &param)
	if err != nil {
		h.handleError(ctx, w, http.StatusBadRequest, "handler.GetJSON: %s", err.Error())
		return
	}

	// Validation
	v := validator.New()
	if err := v.Struct(param); err != nil {
		h.handleError(ctx, w, http.StatusBadRequest, "v.Struct: %s", err.Error())
		return
	}

	err = h.Svc.MessageByUserID(ctx, param.AppID, param.UserID, param.Message)
	if err != nil {
		h.handleError(ctx, w, http.StatusBadRequest, "h.Svc.MessageByUserID: %s", err.Error())
		return
	}
	handler.RenderSuccess(w)
}

// SendToken ... Tokenでプッシュ通知を送信する
func (h *SendHandler) SendToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param model.TaskQueueParamSendToken
	err := handler.GetJSON(r, &param)
	if err != nil {
		h.handleError(ctx, w, http.StatusBadRequest, "handler.GetJSON: %s", err.Error())
		return
	}

	// Validation
	v := validator.New()
	if err := v.Struct(param); err != nil {
		h.handleError(ctx, w, http.StatusBadRequest, "v.Struct: %s", err.Error())
		return
	}

	err = h.Svc.MessageByToken(ctx, param.Token, param.Message)
	if err != nil {
		h.handleError(ctx, w, http.StatusBadRequest, "h.Svc.MessageByToken: %s", err.Error())
		return
	}
	handler.RenderSuccess(w)
}

func (h *SendHandler) handleError(ctx context.Context, w http.ResponseWriter, status int, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Warningf(ctx, msg)
	handler.RenderError(w, status, msg)
}

// NewSendHandler ... SendHandlerを作成する
func NewSendHandler(svc service.Sender) *SendHandler {
	return &SendHandler{
		Svc: svc,
	}
}

package worker

import (
	"net/http"

	"github.com/rabee-inc/push/src/handler"
	"github.com/rabee-inc/push/src/lib/log"
	"github.com/rabee-inc/push/src/model"
	"github.com/rabee-inc/push/src/service"
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
		log.Warningm(ctx, "handler.GetJSON", err)
		handler.HandleError(ctx, w, http.StatusBadRequest, "handler.GetJSON: %s", err.Error())
		return
	}

	err = h.Svc.MessageByUserIDs(ctx, param.UserIDs, param.Message)
	if err != nil {
		log.Warningm(ctx, "h.Svc.MessageByUserIDs", err)
		handler.HandleError(ctx, w, http.StatusBadRequest, "h.Svc.MessageByUserIDs: %s", err.Error())
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
		log.Warningm(ctx, "handler.GetJSON", err)
		handler.HandleError(ctx, w, http.StatusBadRequest, "handler.GetJSON: %s", err.Error())
		return
	}
	err = h.Svc.MessageByUserID(ctx, param.UserID, param.Message)
	if err != nil {
		log.Warningm(ctx, "h.Svc.MessageByUserID", err)
		handler.HandleError(ctx, w, http.StatusBadRequest, "h.Svc.MessageByUserID: %s", err.Error())
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
		log.Warningm(ctx, "handler.GetJSON", err)
		handler.HandleError(ctx, w, http.StatusBadRequest, "handler.GetJSON: %s", err.Error())
		return
	}
	err = h.Svc.MessageByToken(ctx, param.Token, param.Message)
	if err != nil {
		log.Warningm(ctx, "h.Svc.MessageByToken", err)
		handler.HandleError(ctx, w, http.StatusBadRequest, "h.Svc.MessageByToken: %s", err.Error())
		return
	}
	handler.RenderSuccess(w)
}

// NewSendHandler ... SendHandlerを作成する
func NewSendHandler(svc service.Sender) *SendHandler {
	return &SendHandler{
		Svc: svc,
	}
}

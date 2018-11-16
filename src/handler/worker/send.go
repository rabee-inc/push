package worker

import (
	"net/http"

	"github.com/aikizoku/push/src/handler"
	"github.com/aikizoku/push/src/model"
	"github.com/aikizoku/push/src/service"
)

// SendHandler ... サンプルのハンドラ定義
type SendHandler struct {
	Svc service.Sender
}

// SendUserIDs ... UserIDsからUserIDに分割してプッシュ通知を送信する
func (h *SendHandler) SendUserIDs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var param model.SendUserIDs
	err := handler.GetJSON(r, &param)
	if err != nil {
		handler.HandleError(ctx, w, http.StatusBadRequest, "handler.GetJSON: %s", err.Error())
		return
	}

	err = h.Svc.SendMessageToUserIDs(ctx, param.UserIDs, param.Message)
	if err != nil {
		handler.HandleError(ctx, w, http.StatusBadRequest, "h.Svc.SendMessageToUserIDs: %s", err.Error())
		return
	}
	handler.RenderSuccess(w)
}

// SendUserID ... UserIDからTokenを引いてプッシュ通知を送信する
func (h *SendHandler) SendUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var param model.SendUserID
	err := handler.GetJSON(r, &param)
	if err != nil {
		handler.HandleError(ctx, w, http.StatusBadRequest, "handler.GetJSON: %s", err.Error())
		return
	}
	err = h.Svc.SendMessageToUserID(ctx, param.UserID, param.Message)
	if err != nil {
		handler.HandleError(ctx, w, http.StatusBadRequest, "h.Svc.SendMessageToUserID: %s", err.Error())
		return
	}
	handler.RenderSuccess(w)
}

// SendToken ... Tokenでプッシュ通知を送信する
func (h *SendHandler) SendToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var param model.SendToken
	err := handler.GetJSON(r, &param)
	if err != nil {
		handler.HandleError(ctx, w, http.StatusBadRequest, "handler.GetJSON: %s", err.Error())
		return
	}
	err = h.Svc.SendMessageToToken(ctx, param.Token, param.Message)
	if err != nil {
		handler.HandleError(ctx, w, http.StatusBadRequest, "h.Svc.SendMessageToToken: %s", err.Error())
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

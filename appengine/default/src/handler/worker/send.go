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
}

// SendUserIDs ... UserIDsからUserIDに分割してプッシュ通知を送信する
func (h *SendHandler) SendUserIDs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param model.CloudTasksParamSendUserIDs
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

	err = h.sSvc.MessageByUserIDs(ctx, param.AppID, param.UserIDs, param.Message)
	if err != nil {
		renderer.HandleError(ctx, w, "h.sSvc.MessageByUserIDs", err)
		return
	}
	renderer.Success(ctx, w)
}

// SendUserID ... UserIDからTokenを引いてプッシュ通知を送信する
func (h *SendHandler) SendUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param model.CloudTasksParamSendUserID
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

	err = h.sSvc.MessageByUserID(ctx, param.AppID, param.UserID, param.Message)
	if err != nil {
		renderer.HandleError(ctx, w, "h.sSvc.MessageByUserID", err)
		return
	}
	renderer.Success(ctx, w)
}

// NewSendHandler ... SendHandlerを作成する
func NewSendHandler(sSvc service.Sender) *SendHandler {
	return &SendHandler{
		sSvc: sSvc,
	}
}

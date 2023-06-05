package worker

import (
	"net/http"

	"github.com/rabee-inc/go-pkg/errcode"
	"github.com/rabee-inc/go-pkg/parameter"
	"github.com/rabee-inc/go-pkg/renderer"
	"github.com/rabee-inc/push/appengine/push/src/handler"
	"github.com/rabee-inc/push/appengine/push/src/model/input"
	"github.com/rabee-inc/push/appengine/push/src/service"
)

type Sender struct {
	sSender service.Sender
}

func NewSender(
	sSender service.Sender,
) *Sender {
	return &Sender{
		sSender,
	}
}

func (h *Sender) Users(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param input.WorkerSendUsers
	err := parameter.GetJSON(r, &param)
	if err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, err)
		return
	}

	if err = handler.Validate(param); err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, err)
		return
	}

	err = h.sSender.Users(ctx, &param)
	if err != nil {
		renderer.HandleError(ctx, w, err)
		return
	}
	renderer.Success(ctx, w)
}

func (h *Sender) User(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param input.WorkerSendUser
	err := parameter.GetJSON(r, &param)
	if err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, err)
		return
	}

	if err = handler.Validate(param); err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, err)
		return
	}

	err = h.sSender.User(ctx, &param)
	if err != nil {
		renderer.HandleError(ctx, w, err)
		return
	}
	renderer.Success(ctx, w)
}

func (h *Sender) Reserved(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param service.SenderReservedInput
	err := parameter.GetForms(ctx, r, &param)
	if err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, err)
		return
	}

	if err = handler.Validate(param); err != nil {
		err = errcode.Set(err, http.StatusBadRequest)
		renderer.HandleError(ctx, w, err)
		return
	}

	err = h.sSender.Reserved(ctx, &param)
	if err != nil {
		renderer.HandleError(ctx, w, err)
		return
	}
	renderer.Success(ctx, w)
}

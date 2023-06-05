package handler

import (
	"net/http"

	"github.com/rabee-inc/go-pkg/log"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("pong"))
	if err != nil {
		panic(err)
	}
	log.SetResponseStatus(ctx, http.StatusOK)
}

package app

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/rabee-inc/push/appengine/default/src/handler"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	registActions(d)

	r.Use(d.Log.Handle)

	r.Get("/ping", handler.Ping)

	// API
	r.Route("/api", func(r chi.Router) {
		r.With(d.InternalAuth.Handle).Post("/rpc", d.JSONRPC2Handler.Handle)
	})

	// Worker
	r.Route("/worker", func(r chi.Router) {
		r.With(d.InternalAuth.Handle).Post("/send/users", d.SendHandler.SendUserIDs)
		r.With(d.InternalAuth.Handle).Post("/send/user", d.SendHandler.SendUserID)
		r.With(d.InternalAuth.Handle).Post("/send/token", d.SendHandler.SendToken)
	})

	http.Handle("/", r)
}

func registActions(d *Dependency) {
	d.JSONRPC2Handler.Register("entry", d.EntryAction)
	d.JSONRPC2Handler.Register("leave", d.LeaveAction)
	d.JSONRPC2Handler.Register("send", d.SendAction)
}

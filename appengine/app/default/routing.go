package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rabee-inc/push/src/handler"
	"github.com/rabee-inc/push/src/lib/log"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	registActions(d)

	// ログ
	r.Use(log.Handle)

	// API
	r.Route("/api", func(r chi.Router) {
		// 内部認証
		r.Use(d.InternalAuth.Handle)

		// JSONRPC2
		r.Post("/rpc", d.JSONRPC2Handler.Handle)
	})

	// Worker
	r.Route("/worker", func(r chi.Router) {
		// 内部認証
		r.Use(d.InternalAuth.Handle)

		// Task
		r.Post("/send/users", d.SendHandler.SendUserIDs)
		r.Post("/send/user", d.SendHandler.SendUserID)
		r.Post("/send/token", d.SendHandler.SendToken)
	})

	r.Get("/ping", handler.Ping)

	http.Handle("/", r)
}

func registActions(d *Dependency) {
	d.JSONRPC2Handler.Register("entry", d.EntryAction)
	d.JSONRPC2Handler.Register("leave", d.LeaveAction)
	d.JSONRPC2Handler.Register("send", d.SendAction)
}
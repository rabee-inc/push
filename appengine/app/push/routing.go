package main

import (
	"net/http"

	"github.com/aikizoku/push/src/handler"
	"github.com/aikizoku/push/src/lib/log"
	"github.com/go-chi/chi"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	// ログ
	r.Use(log.Handle)

	// API
	r.Route("/api", func(r chi.Router) {
		// 内部認証
		r.Use(d.InternalAuth.Handle)

		// JSONRPC2
		r.Route("/rpc", func(r chi.Router) {
			r.Use(d.JSONRPC2.Handle)
			r.Post("/", handler.Empty)
		})
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

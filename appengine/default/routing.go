package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/rabee-inc/push/appengine/src/handler"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	registActions(d)

	// ログをリクエスト単位でまとめるため、情報をContextに保持する
	r.Use(d.Log.Handle)

	// 障害検知でサーバーの生存確認のため、pingリクエストを用意する
	r.Get("/ping", handler.Ping)

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

	http.Handle("/", r)
}

func registActions(d *Dependency) {
	d.JSONRPC2Handler.Register("entry", d.EntryAction)
	d.JSONRPC2Handler.Register("leave", d.LeaveAction)
	d.JSONRPC2Handler.Register("send", d.SendAction)
}

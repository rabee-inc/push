package main

import (
	"net/http"

	"github.com/aikizoku/push/src/handler"
	"github.com/go-chi/chi"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {

	r.Route("/api", func(r chi.Router) {
		r.Route("/rpc", func(r chi.Router) {
			r.Use(d.JSONRPC2.Handle)
			r.Post("/", handler.Empty)
		})
	})

	r.Get("/ping", handler.Ping)

	http.Handle("/", r)
}

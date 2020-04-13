package app

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/rabee-inc/push/appengine/push/src/handler"
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
		// Tasks
		r.With(d.InternalAuth.Handle).Post("/send/users", d.SendHandler.SendByUsers)
		r.With(d.InternalAuth.Handle).Post("/send/user", d.SendHandler.SendByUser)

		// Scheduler
		r.With(d.InternalAuth.Handle).Get("/send/reserved", d.SendHandler.SendByReserved)
	})

	http.Handle("/", r)
}

func registActions(d *Dependency) {
	d.JSONRPC2Handler.Register("entry", d.EntryAction)
	d.JSONRPC2Handler.Register("leave", d.LeaveAction)
	d.JSONRPC2Handler.Register("send_by_users", d.SendByUsersAction)
	d.JSONRPC2Handler.Register("send_by_all_users", d.SendByAllUsersAction)
	d.JSONRPC2Handler.Register("get_reserve", d.GetReserve)
	d.JSONRPC2Handler.Register("list_reserve", d.ListReserve)
	d.JSONRPC2Handler.Register("create_reserve", d.CreateReserve)
	d.JSONRPC2Handler.Register("update_reserve", d.UpdateReserve)
}

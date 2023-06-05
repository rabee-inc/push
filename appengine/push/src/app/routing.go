package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rabee-inc/push/appengine/push/src/handler"
)

func Routing(r *chi.Mux, d *Dependency) {
	setActions(d)

	r.Use(d.Log.Handle)

	r.Get("/ping", handler.Ping)

	// API
	r.Route("/api", func(r chi.Router) {
		r.With(d.InternalAuth.Handle).Post("/rpc", d.JSONRPC2Handler.Handle)
	})

	// Worker
	r.Route("/worker", func(r chi.Router) {
		// Tasks
		r.With(d.InternalAuth.Handle).Post("/send/users", d.Sender.Users)
		r.With(d.InternalAuth.Handle).Post("/send/user", d.Sender.User)

		// Scheduler
		r.With(d.InternalAuth.Handle).Get("/send/reserved", d.Sender.Reserved)
	})

	http.Handle("/", r)
}

func setActions(d *Dependency) {
	d.JSONRPC2Handler.Register("entry", d.RegisterEntry)
	d.JSONRPC2Handler.Register("leave", d.RegisterLeave)
	d.JSONRPC2Handler.Register("send_by_all_users", d.SenderAllUsers)
	d.JSONRPC2Handler.Register("send_by_users", d.SenderUsers)
	d.JSONRPC2Handler.Register("get_reserve", d.ReserveGet)
	d.JSONRPC2Handler.Register("list_reserve", d.ReserveList)
	d.JSONRPC2Handler.Register("create_reserve", d.ReserveCreate)
	d.JSONRPC2Handler.Register("update_reserve", d.ReserveUpdate)
}

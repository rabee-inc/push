package main

import (
	"github.com/aikizoku/push/src/handler/api"
	"github.com/aikizoku/push/src/handler/worker"
	"github.com/aikizoku/push/src/lib/internalauth"
	"github.com/aikizoku/push/src/lib/jsonrpc2"
	"github.com/aikizoku/push/src/repository"
	"github.com/aikizoku/push/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	InternalAuth *internalauth.Middleware
	JSONRPC2     *jsonrpc2.Middleware
	SendHandler  *worker.SendHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Config
	iaToken := internalauth.GetToken()

	// Repository
	tRepo := repository.NewTokenDatastore()
	fRepo := repository.NewFcm()

	// Service
	eSvc := service.NewEntry(tRepo)
	sSvc := service.NewSender(tRepo, fRepo)

	// Middleware
	d.InternalAuth = internalauth.NewMiddleware(iaToken)
	d.JSONRPC2 = jsonrpc2.NewMiddleware()
	d.JSONRPC2.Register("entry", api.NewEntryHandler(eSvc))
	d.JSONRPC2.Register("send", api.NewSendHandler(sSvc))

	// Handler
	d.SendHandler = worker.NewSendHandler(sSvc)
}

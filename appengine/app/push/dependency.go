package main

import (
	"github.com/rabee-inc/push/src/handler/api"
	"github.com/rabee-inc/push/src/handler/worker"
	"github.com/rabee-inc/push/src/lib/internalauth"
	"github.com/rabee-inc/push/src/lib/jsonrpc2"
	"github.com/rabee-inc/push/src/repository"
	"github.com/rabee-inc/push/src/service"
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
	tRepo := repository.NewTokenFirestore()
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

package main

import (
	"os"

	"github.com/rabee-inc/push/src/lib/cloudfirestore"

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
	crePath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if crePath == "" {
		panic("no config GOOGLE_APPLICATION_CREDENTIALS")
	}
	svrKey := os.Getenv("FCM_SERVER_KEY")
	if svrKey == "" {
		panic("no config FCM_SERVER_KEY")
	}

	// Client
	fCli, err := cloudfirestore.NewClient(crePath)
	if err != nil {
		panic(err.Error())
	}

	// Repository
	tRepo := repository.NewTokenFirestore(fCli)
	fRepo := repository.NewFcm(svrKey)

	// Service
	eSvc := service.NewEntry(tRepo)
	sSvc := service.NewSender(tRepo, fRepo)

	// Middleware
	d.InternalAuth = internalauth.NewMiddleware(iaToken)

	// JSONRPC2
	d.JSONRPC2 = jsonrpc2.NewMiddleware()
	d.JSONRPC2.Register("entry", api.NewEntryHandler(eSvc))
	d.JSONRPC2.Register("send", api.NewSendHandler(sSvc))

	// Handler
	d.SendHandler = worker.NewSendHandler(sSvc)
}

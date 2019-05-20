package main

import (
	"os"

	"firebase.google.com/go/messaging"
	"github.com/rabee-inc/push/src/config"
	"github.com/rabee-inc/push/src/handler/api"
	"github.com/rabee-inc/push/src/handler/worker"
	"github.com/rabee-inc/push/src/lib/cloudfirestore"
	"github.com/rabee-inc/push/src/lib/internalauth"
	"github.com/rabee-inc/push/src/lib/jsonrpc2"
	"github.com/rabee-inc/push/src/repository"
	"github.com/rabee-inc/push/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	InternalAuth    *internalauth.Middleware
	JSONRPC2Handler *jsonrpc2.Handler
	EntryAction     *api.EntryAction
	LeaveAction     *api.LeaveAction
	SendAction      *api.SendAction
	SendHandler     *worker.SendHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Config
	iaToken := internalauth.GetToken()
	crePath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if crePath == "" {
		panic("no config GOOGLE_APPLICATION_CREDENTIALS")
	}

	// FCM
	var env string
	appIDs := config.GetFCMAppIDs()
	fcmClients := map[string]*messaging.Client{}
	fcmServerKeys := map[string]string{}
	if config.IsEnvProduction() {
		env = "production"
	} else {
		env = "staging"
	}
	for _, appID := range appIDs {
		fcmEnv := config.GetFCMEnv(env, appID)
		fcmClients[appID] = config.GetClient(env, appID)
		fcmServerKeys[appID] = fcmEnv.ServerKey
	}

	// Repository(Firestore)
	fCli := cloudfirestore.NewClient(crePath)
	tRepo := repository.NewTokenFirestore(fCli)

	/*
		// Repository(MySQL)
		mCfg := mysql.NewConfig("push")
		mCli := mysql.NewClient(mCfg)
		tRepo := repository.NewTokenMySQL(mCli)
	*/

	/*
		// Repository(Dummy)
		tRepo := repository.NewTokenDummy()
	*/

	// Repository
	fRepo := repository.NewFcm(fcmClients, fcmServerKeys)

	// Service
	rSvc := service.NewRegister(tRepo)
	sSvc := service.NewSender(tRepo, fRepo)

	// Middleware
	d.InternalAuth = internalauth.NewMiddleware(iaToken)

	// Action
	d.EntryAction = api.NewEntryAction(rSvc)
	d.LeaveAction = api.NewLeaveAction(rSvc)
	d.SendAction = api.NewSendAction(sSvc)

	// Handler
	d.JSONRPC2Handler = jsonrpc2.NewHandler()
	d.SendHandler = worker.NewSendHandler(sSvc)
}

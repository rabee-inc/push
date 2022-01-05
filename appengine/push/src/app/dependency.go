package app

import (
	"github.com/rabee-inc/go-pkg/cloudfirestore"
	"github.com/rabee-inc/go-pkg/cloudtasks"
	"github.com/rabee-inc/go-pkg/deploy"
	"github.com/rabee-inc/go-pkg/internalauth"
	"github.com/rabee-inc/go-pkg/jsonrpc2"
	"github.com/rabee-inc/go-pkg/log"

	"github.com/rabee-inc/push/appengine/push/src/config"
	"github.com/rabee-inc/push/appengine/push/src/handler/api"
	"github.com/rabee-inc/push/appengine/push/src/handler/worker"
	"github.com/rabee-inc/push/appengine/push/src/repository"
	"github.com/rabee-inc/push/appengine/push/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	Log                  *log.Middleware
	InternalAuth         *internalauth.Middleware
	JSONRPC2Handler      *jsonrpc2.Handler
	EntryAction          *api.EntryAction
	LeaveAction          *api.LeaveAction
	SendByUsersAction    *api.SendByUsersAction
	SendByAllUsersAction *api.SendByAllUsersAction
	GetReserve           *api.GetReserveAction
	ListReserve          *api.ListReserveAction
	CreateReserve        *api.CreateReserveAction
	UpdateReserve        *api.UpdateReserveAction
	SendHandler          *worker.SendHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject(e *Environment) {
	// Client
	var lCli log.Writer
	if deploy.IsLocal() {
		lCli = log.NewWriterStdout()
	} else {
		lCli = log.NewWriterStackdriver(e.ProjectID)
	}
	fCli := cloudfirestore.NewClient(e.ProjectID)
	tCli := cloudtasks.NewClient(e.Port, e.Deploy, e.ProjectID, "push", e.LocationID, e.InternalAuthToken)
	fcmCli := config.GetClient(e.ProjectID)

	// Repository
	tRepo := repository.NewToken(fCli)
	fRepo := repository.NewFcm(fcmCli, e.FCMServerKey)
	rRepo := repository.NewReserve(fCli)

	// Service
	rgSvc := service.NewRegister(tRepo, fRepo)
	sSvc := service.NewSender(tRepo, fRepo, rRepo, tCli, fCli)
	rSvc := service.NewReserve(rRepo, fCli)

	// Middleware
	d.Log = log.NewMiddleware(lCli, e.MinLogSeverity)
	d.InternalAuth = internalauth.NewMiddleware(e.InternalAuthToken)

	// Handler
	d.JSONRPC2Handler = jsonrpc2.NewHandler()
	d.SendHandler = worker.NewSendHandler(sSvc, rSvc)

	// Action
	d.EntryAction = api.NewEntryAction(rgSvc)
	d.LeaveAction = api.NewLeaveAction(rgSvc)
	d.SendByUsersAction = api.NewSendByUsersAction(tCli)
	d.SendByAllUsersAction = api.NewSendByAllUsersAction(sSvc)
	d.GetReserve = api.NewGetReserveAction(rSvc)
	d.ListReserve = api.NewListReserveAction(rSvc)
	d.CreateReserve = api.NewCreateReserveAction(rSvc)
	d.UpdateReserve = api.NewUpdateReserveAction(rSvc)
}

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

type Dependency struct {
	Log             *log.Middleware
	InternalAuth    *internalauth.Middleware
	JSONRPC2Handler *jsonrpc2.Handler
	RegisterEntry   *api.RegisterEntry
	RegisterLeave   *api.RegisterLeave
	SenderAllUsers  *api.SenderAllUsers
	SenderUsers     *api.SenderUsers
	ReserveGet      *api.ReserveGet
	ReserveList     *api.ReserveList
	ReserveCreate   *api.ReserveCreate
	ReserveUpdate   *api.ReserveUpdate
	Sender          *worker.Sender
}

func (d *Dependency) Inject(e *Environment) {
	// Client
	var cLog log.Writer
	if deploy.IsLocal() {
		cLog = log.NewWriterStdout()
	} else {
		cLog = log.NewWriterStackdriver(e.ProjectID)
	}
	cFirestore := cloudfirestore.NewClient(e.ProjectID)
	cTasks := cloudtasks.NewClient(e.Port, e.Deploy, e.ProjectID, "push", e.LocationID, e.InternalAuthToken)
	fcmCli := config.GetClient(e.ProjectID)

	// Repository
	rToken := repository.NewToken(cFirestore)
	rFCM := repository.NewFCM(fcmCli, e.FCMServerKey)
	rReserve := repository.NewReserve(cFirestore)

	// Service
	rgSvc := service.NewRegister(rToken, rFCM)
	sSvc := service.NewSender(rToken, rFCM, rReserve, cTasks, cFirestore)
	rSvc := service.NewReserve(rReserve, cFirestore)

	// Middleware
	d.Log = log.NewMiddleware(cLog, e.MinLogSeverity)
	d.InternalAuth = internalauth.NewMiddleware(e.InternalAuthToken)

	// Handler
	d.JSONRPC2Handler = jsonrpc2.NewHandler()
	d.Sender = worker.NewSender(sSvc)

	// Action
	d.RegisterEntry = api.NewRegisterEntry(rgSvc)
	d.RegisterLeave = api.NewRegisterLeave(rgSvc)
	d.SenderAllUsers = api.NewSenderAllUsers(sSvc)
	d.SenderUsers = api.NewSenderUsers(cTasks)
	d.ReserveGet = api.NewReserveGet(rSvc)
	d.ReserveList = api.NewReserveList(rSvc)
	d.ReserveCreate = api.NewReserveCreate(rSvc)
	d.ReserveUpdate = api.NewReserveUpdate(rSvc)
}

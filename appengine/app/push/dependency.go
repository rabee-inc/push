package main

import (
	"github.com/aikizoku/push/src/handler/api"
	"github.com/aikizoku/push/src/lib/jsonrpc2"
	"github.com/aikizoku/push/src/repository"
	"github.com/aikizoku/push/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	JSONRPC2 *jsonrpc2.Middleware
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Repository
	tRepo := repository.NewTokenFirestore()

	// Service
	eSvc := service.NewEntry(tRepo)

	// Middleware
	d.JSONRPC2 = jsonrpc2.NewMiddleware()
	d.JSONRPC2.Register("entry", api.NewEntryHandler(eSvc))

	// Handler
}

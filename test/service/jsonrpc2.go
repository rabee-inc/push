package service

import "github.com/rabee-inc/push/test/model"

// JSONRPC2 ... JSONRPC2形式のAPI通信を行う
type JSONRPC2 interface {
	Send(name string, method string, params map[string]interface{})
	GetAPIs() []*model.API
}

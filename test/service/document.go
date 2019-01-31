package service

import "github.com/rabee-inc/push/test/model"

// Document ... ドキュメントを操作する
type Document interface {
	RemoveAll()
	Distributes(tmplPath string, apis []*model.API)
}

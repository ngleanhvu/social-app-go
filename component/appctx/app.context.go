package appctx

import (
	"crud-go/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appContext struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, uploadprovider uploadprovider.UploadProvider) *appContext {
	return &appContext{db: db, uploadProvider: uploadprovider}
}

func (ctx *appContext) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appContext) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

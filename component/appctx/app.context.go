package appctx

import (
	"crud-go/component/uploadprovider"
	"crud-go/pubsub"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	GetPubSub() pubsub.PubSub
}

type appContext struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	pubSub         pubsub.PubSub
}

func NewAppContext(db *gorm.DB,
	uploadprovider uploadprovider.UploadProvider,
	pubsub pubsub.PubSub) *appContext {
	return &appContext{db: db, uploadProvider: uploadprovider, pubSub: pubsub}
}

func (ctx *appContext) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appContext) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appContext) GetPubSub() pubsub.PubSub {
	return ctx.pubSub
}

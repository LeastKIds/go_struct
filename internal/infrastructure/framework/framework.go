package framework

import (
	"net/http"

	"gorm.io/gorm"
)

type IFramework interface {
	Use(middleware func(http.Handler) http.Handler)
	Group(prefix string) IGroup
	Start(address string) error
}

type IGroup interface {
	GET(path string, handler HandlerFuncWithDB)
	POST(path string, handler HandlerFuncWithDB)
	PUT(path string, handler HandlerFuncWithDB)
	DELETE(path string, handler HandlerFuncWithDB)

	Group(prefix string) IGroup
}

func NewFramework(db *gorm.DB) IFramework {
	return NewEchoFramework(db)
}

type HandlerFuncWithDB func(http.ResponseWriter, *http.Request, *gorm.DB)

type InterceptorFunc func(HandlerFuncWithDB) HandlerFuncWithDB

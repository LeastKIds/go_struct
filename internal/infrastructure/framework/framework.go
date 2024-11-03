package framework

import (
	"net/http"
)

type IFramework interface {
	Use(middleware func(http.Handler) http.Handler)
	Group(prefix string) IGroup
	Start(address string) error
}

type IGroup interface {
	GET(path string, handler func(http.ResponseWriter, *http.Request))
	POST(path string, handler func(http.ResponseWriter, *http.Request))
	PUT(path string, handler func(http.ResponseWriter, *http.Request))
	DELETE(path string, handler func(http.ResponseWriter, *http.Request))

	Group(prefix string) IGroup
}

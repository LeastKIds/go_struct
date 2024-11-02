package handler

import "github.com/LeastKIds/go_struct/internal/adapter/handler/dummy"

type Handler struct {
	Dummy dummy.IDummyHandler
}

func NewHandler(dummy dummy.IDummyHandler) *Handler {
	return &Handler{Dummy: dummy}
}

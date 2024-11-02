package router

type IRouter interface{}

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

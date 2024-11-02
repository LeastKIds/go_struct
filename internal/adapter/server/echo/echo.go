package echo

import "github.com/labstack/echo/v4"

type EchoServer struct {
	e *echo.Echo
}

func NewEchoServer() *EchoServer {
	return &EchoServer{
		e: echo.New(),
	}
}

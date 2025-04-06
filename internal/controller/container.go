package controller

import "github.com/phat-ngoc-anh/backend/internal/controller/http"

type ApiContainer struct {
	HttpServer *http.Server
}

func NewApiContainer(httpServer *http.Server) *ApiContainer {
	return &ApiContainer{HttpServer: httpServer}
}

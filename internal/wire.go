//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/phat-ngoc-anh/backend/internal/controller"
	"github.com/phat-ngoc-anh/backend/internal/controller/http"
	v1 "github.com/phat-ngoc-anh/backend/internal/controller/http/v1"
	"github.com/phat-ngoc-anh/backend/internal/database"
	repositoryimplement "github.com/phat-ngoc-anh/backend/internal/repository/implement"
	serviceimplement "github.com/phat-ngoc-anh/backend/internal/service/implement"
)

var container = wire.NewSet(
	controller.NewApiContainer,
)

// may have grpc server in the future
var serverSet = wire.NewSet(
	http.NewServer,
)

// handler === controller | with service and repository layers to form 3 layers architecture
var handlerSet = wire.NewSet(
	v1.NewStudentHandler,
	v1.NewInvoiceHandler,
)

var serviceSet = wire.NewSet(
	serviceimplement.NewStudentService,
	serviceimplement.NewInvoiceService,
)

var repositorySet = wire.NewSet(
	repositoryimplement.NewStudentRepository,
)

func InitializeContainer(
	db database.Db,
) *controller.ApiContainer {
	wire.Build(serverSet, handlerSet, serviceSet, repositorySet, container)
	return &controller.ApiContainer{}
}

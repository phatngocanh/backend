package startup

import (
	"github.com/phat-ngoc-anh/backend/internal"
	"github.com/phat-ngoc-anh/backend/internal/controller"
	"github.com/phat-ngoc-anh/backend/internal/database"
)

func registerDependencies() *controller.ApiContainer {
	// Open database connection
	db := database.Open()

	return internal.InitializeContainer(db)
}

func Execute() {
	container := registerDependencies()
	container.HttpServer.Run()
}

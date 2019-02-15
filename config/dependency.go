package config

import (
	"time"

	hostRest "github.com/lexicforlxd/backend-reloaded/host/delivery/rest"
	hostRepo "github.com/lexicforlxd/backend-reloaded/host/repository"
	hostUsecase "github.com/lexicforlxd/backend-reloaded/host/usecase"
	"github.com/lexicforlxd/backend-reloaded/models"
	"github.com/lexicforlxd/backend-reloaded/resolvers"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

// BuildContainer builds the container for DI
func BuildContainer() *dig.Container {
	container := dig.New()
	// Database
	container.Provide(models.CreateConnection)

	// Host
	container.Provide(hostRepo.NewHostRepository)
	container.Provide(hostUsecase.NewHostUsecase)
	container.Provide(hostRest.NewHostHandler)

	// GraphQL
	container.Provide(resolvers.NewResolver)

	container.Provide(func() time.Duration {
		return time.Duration(viper.GetInt("timeout")) * time.Second
	})

	return container
}

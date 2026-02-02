package hitrixapp

import (
	"awesomeProject1/internal/infrastructure/persistence/beeorm"

	"github.com/coretrix/hitrix"
	"github.com/coretrix/hitrix/service/component/app"
	"github.com/coretrix/hitrix/service/registry"
)

func New() (*hitrix.Hitrix, func()) {
	r := hitrix.New(
		"TodoSimpleProject",
		"your-secret",
	)

	r.RegisterRedisPools(&app.RedisPools{Persistent: beeorm.RedisPool})

	r.RegisterDIGlobalService(
		registry.ServiceProviderErrorLogger(),
		registry.ServiceProviderConfigDirectory("./config"),
		registry.ServiceProviderOrmRegistry(beeorm.Init),
		registry.ServiceProviderOrmEngine(),
	)

	a, cleanup := r.Build()
	return a, cleanup
}

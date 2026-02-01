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

	r.RegisterDIGlobalService(
		registry.ServiceProviderErrorLogger(),
		registry.ServiceProviderConfigDirectory("./config"),
		registry.ServiceProviderOrmRegistry(beeorm.Init),
		registry.ServiceProviderOrmEngine(),
	).RegisterRedisPools(&app.RedisPools{Persistent: "your pool here"})

	a, cleanup := r.Build()
	return a, cleanup
}

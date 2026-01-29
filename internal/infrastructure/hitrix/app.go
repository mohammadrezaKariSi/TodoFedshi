package hitrixapp

import "github.com/coretrix/hitrix"

func New() (*hitrix.Hitrix, func()) {
	registry := hitrix.New(
		"TodoSimpleProject",
		"your-secret",
	)

	// Example (optional):
	// registry.RegisterConfigPath("./configs")

	app, cleanup := registry.Build()
	return app, cleanup
}

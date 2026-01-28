package hitrix

import "github.com/coretrix/hitrix"

func New() *hitrix.App {
	app := hitrix.NewApp()

	// register entities
	app.RegisterEntity(&beeorm.UserEntity{})

	return app
}

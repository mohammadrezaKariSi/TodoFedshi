package scripts

import (
	"context"
	"fmt"
	"time"

	"github.com/coretrix/hitrix/service/component/app"
	"github.com/latolukasz/beeorm"
)

type TodoScript struct{}

func (t TodoScript) Description() string {
	return "todo-simple-script"
}

func (t TodoScript) Run(ctx context.Context, exit app.IExit, ormService *beeorm.Engine) {
	fmt.Println("script is running")
}

func (t TodoScript) Unique() bool {
	return false
}

func (t TodoScript) Interval() time.Duration {
	return time.Second * 10
}

package consumers

import (
	beeorm2 "awesomeProject1/internal/infrastructure/persistence/beeorm"
	"context"
	"fmt"

	"github.com/coretrix/hitrix/pkg/queue/streams"
	"github.com/latolukasz/beeorm"
)

type DirtyConsumer struct {
	Ctx context.Context
}

func (c *DirtyConsumer) Consume(ormService *beeorm.Engine, events []beeorm.Event) error {
	for _, event := range events {
		e := &beeorm2.ToDoEntity{}
		event.Unserialize(e)

		fmt.Println("todo inserted with ID:", e.ID)
	}

	return nil
}

func NewDirtyConsumer() *DirtyConsumer {
	return &DirtyConsumer{}
}

func (c *DirtyConsumer) GetQueueName() string {
	return beeorm2.ToDoChannel
}

func (c *DirtyConsumer) GetGroupName(suffix *string) string {
	return streams.GetGroupName(c.GetQueueName(), suffix)
}

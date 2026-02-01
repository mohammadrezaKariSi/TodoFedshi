package consumers

import (
	beeorm2 "awesomeProject1/internal/infrastructure/persistence/beeorm"
	"fmt"

	"github.com/coretrix/hitrix/pkg/queue/streams"
	"github.com/latolukasz/beeorm"
)

type DirtyConsumer struct {
}

func (c *DirtyConsumer) Consume(ormService *beeorm.Engine, event beeorm.Event) error {
	e := &beeorm2.ToDoEntity{}
	event.Unserialize(e)

	fmt.Printf("todo inserted with ID: %+v\n", e.ID)

	return nil
}

func NewDirtyConsumer() *DirtyConsumer {
	return &DirtyConsumer{}
}

func (c *DirtyConsumer) GetQueueName() string {
	return "todo_channel"
}

func (c *DirtyConsumer) GetGroupName(suffix *string) string {
	return streams.GetGroupName(c.GetQueueName(), suffix)
}

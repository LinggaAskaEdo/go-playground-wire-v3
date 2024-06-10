package task

import (
	"github.com/go-co-op/gocron/v2"

	"github.com/linggaaskaedo/go-playground-wire-v3/src/handler/scheduler"
)

type SchedulerTaskImpl struct {
	handlers *scheduler.SchedulerHandlerImpl
}

func NewSchedulerTask(
	handlers *scheduler.SchedulerHandlerImpl,
) *SchedulerTaskImpl {
	return &SchedulerTaskImpl{
		handlers: handlers,
	}
}

func (t *SchedulerTaskImpl) Router(r gocron.Scheduler) {
	t.handlers.Scheduler(r)
}

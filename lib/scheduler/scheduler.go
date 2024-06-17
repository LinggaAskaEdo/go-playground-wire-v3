package scheduler

import (
	"context"

	"github.com/go-co-op/gocron/v2"
	"github.com/rs/zerolog/log"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/scheduler/task"
)

type SchedulerImpl struct {
	SchedulerTask *task.SchedulerTaskImpl
	scheduler      gocron.Scheduler
}

func NewScheduler(
	SchedulerTask *task.SchedulerTaskImpl,
) *SchedulerImpl {
	log.Info().Msg("Initialize scheduler...")

	return &SchedulerImpl{
		SchedulerTask: SchedulerTask,
	}
}

func (p *SchedulerImpl) setupTask(sched gocron.Scheduler) {
	p.SchedulerTask.Router(sched)
}

func (p *SchedulerImpl) Start() {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal().Msg("Failed initiate scheduler")
	}

	p.setupTask(s)
	s.Start()
	log.Info().Msg("Scheduler starting...")
}

func (p *SchedulerImpl) Shutdown(ctx context.Context) error {
	if err := p.scheduler.Shutdown(); err != nil {
		return err
	}

	return nil
}

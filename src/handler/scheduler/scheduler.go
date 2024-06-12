package scheduler

import (
	"context"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/rs/zerolog/log"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/config"
	newssvc "github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/service"
)

type SchedulerHandlerImpl struct {
	newssvc.NewsService
}

func NewSchedulerHandler(
	newsService newssvc.NewsService,
) *SchedulerHandlerImpl {
	return &SchedulerHandlerImpl{
		NewsService: newsService,
	}
}

func (s *SchedulerHandlerImpl) Scheduler(t gocron.Scheduler) {
	ctx := context.Background()

	if config.Get().Module.News.Scheduler.GetNewsRSSEnable {
		_, err := t.NewJob(
			gocron.DurationJob(
				time.Duration(config.Get().Module.News.Scheduler.GetNewsRSSDuration)*time.Minute,
			),
			gocron.NewTask(
				func(ctx context.Context) {
					s.GetLatestNewsRSS(ctx)
				},
				ctx,
			),
		)
		if err != nil {
			log.Err(err).Msg("Error GetLatestNewsRSS")
		}
	}

	if config.Get().Module.News.Scheduler.GetNewsIndexEnable {
		_, err := t.NewJob(
			gocron.DurationJob(
				time.Duration(config.Get().Module.News.Scheduler.GetNewsIndexDuration)*time.Minute,
			),
			gocron.NewTask(
				func(ctx context.Context) {
					s.GetLatestNewsIndex(ctx)
				},
				ctx,
			),
		)
		if err != nil {
			log.Err(err).Msg("Error GetLatestNewsIndex")
		}
	}
}

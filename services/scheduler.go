package services

import (
	"time"

	"github.com/go-co-op/gocron/v2"
)

type SchedulerService struct {
	Scheduler gocron.Scheduler
}

func NewSchedulerService() (*SchedulerService, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	s, err := gocron.NewScheduler(gocron.WithLocation(jst))
	if err != nil {
		return nil, err
	}
	s.Start()

	return &SchedulerService{Scheduler: s}, nil
}

// NOTE: ここでは使わない
func (s *SchedulerService) Shutdown() {
	s.Scheduler.Shutdown()
}

func (s *SchedulerService) NewDailyJob(fn func()) (gocron.Job, error) {
	nj, err := s.Scheduler.NewJob(
		gocron.CronJob("0 9 * * 1-5", false),
		gocron.NewTask(fn),
	)
	if err != nil {
		return nil, err
	}

	return nj, nil
}

package cronworker

import (
	"context"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
)

type Worker struct {
	Execute  func(ctx context.Context) error
	Schedule string
	Name     string
}

type TaskManager struct {
	cron *cron.Cron
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		cron: cron.New(),
	}
}

func (tm *TaskManager) AddWorker(ctx context.Context, worker Worker) error {
	wrapFunc := func() {
		const retryAfterErr = time.Second * 5

		for err := worker.Execute(ctx); err != nil; {
			zerolog.Ctx(ctx).Err(err).Str("cron", worker.Name).Msg("executing cron")

			timer := time.NewTimer(retryAfterErr)

			select {
			case <-ctx.Done():
				timer.Stop()

				return
			case <-timer.C:
				err = worker.Execute(ctx)

				timer.Stop()
			}
		}
	}

	if _, err := tm.cron.AddFunc(worker.Schedule, wrapFunc); err != nil {
		return fmt.Errorf("adding func to cron: %w", err)
	}

	return nil
}

func (tm *TaskManager) Start() {
	tm.cron.Start()
}

func (tm *TaskManager) Stop() context.Context {
	return tm.cron.Stop()
}

package worker

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func (w *Worker) Run(ctx context.Context) error {
	c := cron.New()

	_, err := c.AddFunc(cronMonthlyReset, func() {
		err := w.userRepo.ResetGenerationsPerMonthForAll(ctx)
		if err != nil {
			logrus.Errorf("Worker reset generations per month for all users error: %v", err)
		}
	})
	if err != nil {
		return err
	}

	c.Start()
	logrus.Info("Worker started")

	<-ctx.Done()

	c.Stop()

	return nil
}

package worker

import "context"

type userRepo interface {
	ResetGenerationsPerMonthForAll(ctx context.Context) error
}

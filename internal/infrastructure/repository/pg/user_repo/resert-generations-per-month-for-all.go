package user_repo

import "context"

func (r *Repository) ResetGenerationsPerMonthForAll(ctx context.Context) error {
	_, err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).ExecContext(
		ctx,
		queryResetGenerationsPerMonthForAllUsers,
	)
	if err != nil {
		return err
	}

	return nil
}

package user_repo

import "context"

func (r *Repository) LockForUpdate(ctx context.Context, userID int) error {
	_, err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).ExecContext(
		ctx,
		queryLockUserForUpdate,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}

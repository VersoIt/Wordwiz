package user_repo

import (
	"context"
	"wordwiz/internal/domain/model"
)

func (r *Repository) Update(ctx context.Context, user model.User) error {
	_, err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).ExecContext(
		ctx,
		queryUpdateUser,
		user.ID,
		user.TotalRequests,
		user.CreatedAt,
		user.GenerationsPerMonth,
	)
	if err != nil {
		return err
	}

	return nil
}

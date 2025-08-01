package user_repo

import (
	"context"
	"wordwiz/internal/domain/model"
)

func (r *Repository) Create(
	ctx context.Context,
	user model.User,
) error {
	_, err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).ExecContext(
		ctx,
		queryCreateUser,
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

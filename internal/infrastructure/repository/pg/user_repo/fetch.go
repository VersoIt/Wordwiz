package user_repo

import (
	"context"
	"github.com/samber/lo"
	"wordwiz/internal/domain/model/user"
	"wordwiz/internal/infrastructure/repository/pg/user_repo/entity"
)

func (r *Repository) Fetch(ctx context.Context) (user.Users, error) {
	var users entity.Users

	err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).SelectContext(
		ctx,
		&users,
		queryFetchUsers,
	)
	if err != nil {
		return nil, err
	}

	return lo.Map(users, func(user entity.User, _ int) user.User {
		return user.ToDomain()
	}), nil
}

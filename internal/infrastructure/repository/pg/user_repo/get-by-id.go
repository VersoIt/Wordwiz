package user_repo

import (
	"context"
	"wordwiz/internal/domain/model"
	"wordwiz/internal/infrastructure/repository/pg/user_repo/entity"
)

func (r *Repository) GetByID(ctx context.Context, id int) (model.User, error) {
	var u entity.User

	err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).Get(&u, queryGetUserByID)
	if err != nil {
		return model.User{}, err
	}

	return u.ToDomain(), nil
}

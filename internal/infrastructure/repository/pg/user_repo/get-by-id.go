package user_repo

import (
	"context"
	"database/sql"
	"errors"
	"wordwiz/internal/domain/model"
	"wordwiz/internal/domain/model/user"
	"wordwiz/internal/infrastructure/repository/pg/user_repo/entity"
)

func (r *Repository) GetByID(ctx context.Context, id int) (user.User, error) {
	var u entity.User

	err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).Get(&u, queryGetUserByID, id)
	if errors.Is(err, sql.ErrNoRows) {
		return user.User{}, model.ErrUserNotFound
	}

	if err != nil {
		return user.User{}, err
	}

	return u.ToDomain(), nil
}

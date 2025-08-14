package user

import (
	"context"
	"wordwiz/internal/domain/model/user"
)

type Repo interface {
	Create(ctx context.Context, user user.User) error
	GetByID(ctx context.Context, userID int) (user.User, error)
}

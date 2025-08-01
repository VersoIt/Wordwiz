package user

import (
	"context"
	"wordwiz/internal/domain/model"
)

type Repo interface {
	Create(ctx context.Context) error
	GetByID(ctx context.Context, userID int) (model.User, error)
}

package music_smart_generator

import (
	"context"
	"wordwiz/internal/domain/model"
)

type userRepo interface {
	LockForUpdate(ctx context.Context, userID int) error
	GetByID(ctx context.Context, id int) (model.User, error)
	Update(ctx context.Context, u model.User) error
}

type aiClient interface {
	Do(ctx context.Context, request string) (string, error)
}

package user

import (
	"context"
	"errors"
	"wordwiz/internal/domain/model"
)

func (s *Service) TryCreateUser(ctx context.Context, user model.User) (bool, error) {
	_, err := s.repo.GetByID(ctx, user.ID)
	if errors.Is(err, model.ErrUserNotFound) {
		// unique constraint on userID
		err = s.repo.Create(ctx)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	if err != nil {
		return false, err
	}

	return false, nil
}

package music_smart_generator

import (
	"context"
	"wordwiz/internal/domain/model"
)

func (u *UseCase) Generate(ctx context.Context, text string, userID int) (model.Verses, error) {
	var verses model.Verses

	err := u.txManager.Do(ctx, func(ctx context.Context) error {
		err := u.userRepo.LockForUpdate(ctx, userID)
		if err != nil {
			return err
		}

		user, err := u.userRepo.GetByID(ctx, userID)
		if err != nil {
			return err
		}

		rawText, err := u.aiClient.Do(ctx, text)
		if err != nil {
			return err
		}

		rawVerses := model.RawVerses(rawText)

		verses = rawVerses.ToVerses()

		if user.GenerationsPerMonth >= u.cfg.MaxGenerationsPerMonth {
			return model.ErrGenerationsLimitReached
		}

		user.GenerationsPerMonth++
		user.TotalRequests++

		err = u.userRepo.Update(ctx, user)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return verses, nil
}

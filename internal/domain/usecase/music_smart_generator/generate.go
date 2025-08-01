package music_smart_generator

import (
	"context"
	_ "embed"
	"fmt"
	"wordwiz/internal/domain/model"
	"wordwiz/internal/domain/model/user"
	"wordwiz/internal/domain/model/verse"
)

//go:embed lyrics_gen_prompt
var lyricsGenPrompt string

func (u *UseCase) GenerateWithStats(
	ctx context.Context,
	text string,
	userID int,
) (user.User, verse.Verses, error) {
	var (
		verses verse.Verses
		stats  user.User
	)

	if len(text) == 0 {
		return user.User{}, verse.Verses{}, model.ErrEmptyArgs
	}

	err := u.txManager.Do(ctx, func(ctx context.Context) error {
		_, err := u.userService.TryCreateUser(ctx, *user.New(userID))
		if err != nil {
			return err
		}

		err = u.userRepo.LockForUpdate(ctx, userID)
		if err != nil {
			return err
		}

		stats, err = u.userRepo.GetByID(ctx, userID)
		if err != nil {
			return err
		}

		rawText, err := u.aiClient.Do(ctx, fmt.Sprintf(lyricsGenPrompt, text))
		if err != nil {
			return err
		}

		rawVerses := verse.RawVerses(rawText)

		verses = rawVerses.ToVerses()

		if stats.GenerationsPerMonth >= u.cfg.MaxGenerationsPerMonth {
			return model.ErrGenerationsLimitReached
		}

		stats.GenerationsPerMonth++
		stats.TotalRequests++

		err = u.userRepo.Update(ctx, stats)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return user.User{}, nil, err
	}

	return stats, verses, nil
}

package music_smart_generator

import (
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"wordwiz/config"
	"wordwiz/internal/domain/service/user"
)

type UseCase struct {
	cfg         config.Config
	userRepo    userRepo
	txManager   *manager.Manager
	userService *user.Service
	aiClient    aiClient
}

func New(
	cfg config.Config,
	userRepo userRepo,
	txManager *manager.Manager,
	userService *user.Service,
	aiClient aiClient,
) *UseCase {
	return &UseCase{
		cfg:         cfg,
		userRepo:    userRepo,
		txManager:   txManager,
		userService: userService,
		aiClient:    aiClient,
	}
}

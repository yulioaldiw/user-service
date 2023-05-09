package drivenadapt

import (
	"context"

	"user-service/internal/adapters/driven/repository"
	"user-service/internal/core/domain"
)

type UserAdapter struct {
	userRepository repository.UserRepository
}

func NewUserAdapter(usrep repository.UserRepository) UserAdapter {
	return UserAdapter{
		userRepository: usrep,
	}
}

func (uadpt UserAdapter) Create(ctx context.Context, user domain.User) error {
	return uadpt.userRepository.Save(ctx, user)
}

package repository

import (
	"context"

	"user-service/internal/core/domain"
)

type UserRepository interface {
	Save(ctx context.Context, sqlUser domain.User) error
}

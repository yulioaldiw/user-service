package drivenport

import (
	"context"

	"user-service/internal/core/domain"
)

type UserDB interface {
	Create(ctx context.Context, user domain.User) error
}

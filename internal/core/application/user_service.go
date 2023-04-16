package application

import (
	"context"
	"time"

	"user-service/internal/core/application/dto"
	"user-service/internal/core/domain"
	outDB "user-service/internal/core/ports/driven"

	"github.com/dany0814/go-hexagonal/pkg/encryption"
	"github.com/dany0814/go-hexagonal/pkg/uidgen"
)

type UserService struct {
	userDB outDB.UserDB
}

func NewUserService(userDB outDB.UserDB) UserService {
	return UserService{
		userDB: userDB,
	}
}

func (usrv UserService) Register(ctx context.Context, user dto.User) (*dto.User, error) {
	id := uidgen.New().New()

	newuser, err := domain.NewUser(id, user.Name, user.Lastname, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	pass, err := encryption.HashAndSalt(user.Password)
	if err != nil {
		return nil, err
	}

	passencrypted, _ := domain.NewUserPassword(pass)

	newuser.Password = passencrypted
	newuser.CreatedAt = time.Now()
	newuser.UpdatedAt = time.Now()

	err = usrv.userDB.Create(ctx, newuser)
	if err != nil {
		return nil, err
	}

	user.ID = id
	return &user, nil
}

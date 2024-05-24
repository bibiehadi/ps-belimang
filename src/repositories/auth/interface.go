package authRepository

import (
	"belimang/src/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository interface {
	Create(user entities.Auth) (entities.Auth, error)
	UsernameIsExist(username string) bool
	FindByUsername(username string) (entities.Auth, error)
}

type authRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *authRepository {
	return &authRepository{db}
}

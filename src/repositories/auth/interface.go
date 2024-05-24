package authRepository

import "github.com/jackc/pgx/v5/pgxpool"

type AuthRepository interface {
}

type authRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *authRepository {
	return &authRepository{db}
}

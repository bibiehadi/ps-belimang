package authRepository

import (
	"belimang/src/entities"
	"context"

	"github.com/jackc/pgx/v5"
)

func (repository *authRepository) FindByUsername(username string) (entities.Auth, error) {
	var auth entities.Auth
	var query string = `SELECT id, username, email, role, password, created_at FROM auths WHERE username = $1`
	err := repository.db.QueryRow(context.Background(), query, username).Scan(&auth.Id, &auth.Username, &auth.Email, &auth.Role, &auth.Password, &auth.CreatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return entities.Auth{}, err
		}
	}

	return auth, err
}

func (repository *authRepository) UsernameIsExist(username string) bool {
	var exist string
	var query string = `SELECT username FROM auths WHERE username = $1 LIMIT 1`
	err := repository.db.QueryRow(context.Background(), query, username).Scan(&exist)

	if err != nil {
		if err == pgx.ErrNoRows {
			return false
		}
	}
	return true
}

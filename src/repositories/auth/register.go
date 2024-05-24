package authRepository

import (
	"belimang/src/entities"
	"context"

	"github.com/google/uuid"
)

func (r *authRepository) Create(user entities.Auth) (entities.Auth, error) {
	var id = uuid.NewString()
	var userId string
	var query string = `INSERT INTO auths (id, username, password, email, role) values ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(context.Background(), query, id, user.Username, user.Password, user.Email, user.Role).Scan(&userId)

	if err != nil {
		return entities.Auth{}, err
	}

	user.Id = userId

	return user, err
}

package entities

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Auth struct {
	Id        string    `json:"id" validate:"required"`
	Username  string    `json:"username" validate:"required,min=5,max=30"`
	Password  string    `json:"password" validate:"required,min=5,max=30"`
	Email     string    `json:"email" validate:"required"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=5,max=30"`
	Password string `json:"password" validate:"required,min=5,max=30"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=5,max=30"`
	Password string `json:"password" validate:"required,min=5,max=30"`
}

type AuthResponse struct {
	Token string `json:"token" validate:"required"`
}

type Role string

const (
	User  Role = "user"
	Admin Role = "admin"
)

type CustomClaims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

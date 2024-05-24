package http

import "github.com/jackc/pgx/v5/pgxpool"

type Http struct {
	DB *pgxpool.Pool
}

type iHttp interface {
	Launch()
}

func New(http *Http) iHttp {
	return http
}

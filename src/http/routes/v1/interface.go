package v1

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type V1Routes struct {
	Echo *echo.Group
	Db   *pgxpool.Pool
}

type iV1Routes interface {
	MountAuth()
	MountMerchant()
	MountMerchantItem()
	MountPurchase()
	MountUpload()
}

func New(v1Routes *V1Routes) iV1Routes {
	return v1Routes
}

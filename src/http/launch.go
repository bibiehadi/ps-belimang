package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (r *Http) Launch() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodHead,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
				http.MethodPost,
			},
		},
	))

	// Mount all routes here
	basePath := "/v1"
	baseUrl := e.Group(basePath)
	baseUrl.GET("", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf("API Base Code for %s", os.Getenv("ENVIRONMENT")))
	})

	// v1 := v1routes.New(
	// 	&v1routes.V1Routes{
	// 		Echo: e.Group(basePath),
	// 		Db:   r.DB,
	// 	},
	// )

	// v1.MountAuth()

	e.Logger.Fatal(e.Start(":8080"))
}

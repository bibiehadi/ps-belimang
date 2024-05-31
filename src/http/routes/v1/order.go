package v1

import (
	orderController "belimang/src/http/controller/order"
	itemRepository "belimang/src/repositories/item"
	merchantRepository "belimang/src/repositories/merchant"
	orderService "belimang/src/services/order"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (i *V1Routes) MountPurchase() {
	g := i.Echo.Group("/users")
	merchantRepository := merchantRepository.New(i.Db)
	itemRepository := itemRepository.New(i.Db)
	orderService := orderService.New(merchantRepository, itemRepository)
	orderController := orderController.New(orderService)

	g.POST("/estimate", orderController.Estimate)

	g.POST("/orders", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf("API Base Code for %s", os.Getenv("ENVIRONMENT")))
	})
	g.GET("/orders", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf("API Base Code for %s", os.Getenv("ENVIRONMENT")))
	})
}

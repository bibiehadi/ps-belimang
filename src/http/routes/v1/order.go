package v1

import (
	orderController "belimang/src/http/controller/order"
	"belimang/src/http/middlewares"
	itemRepository "belimang/src/repositories/item"
	merchantRepository "belimang/src/repositories/merchant"
	orderRepository "belimang/src/repositories/order"
	orderService "belimang/src/services/order"
)

func (i *V1Routes) MountPurchase() {
	g := i.Echo.Group("/users")
	g.Use(middlewares.RequireAuth())
	g.Use(middlewares.AuthWithRole("user"))
	merchantRepository := merchantRepository.New(i.Db)
	itemRepository := itemRepository.New(i.Db)
	orderRepository := orderRepository.New(i.Db)
	orderService := orderService.New(merchantRepository, itemRepository, orderRepository)
	orderController := orderController.New(orderService)

	g.POST("/estimate", orderController.Estimate)

	g.POST("/orders", orderController.Order)
	g.GET("/orders", orderController.FindAll)
}

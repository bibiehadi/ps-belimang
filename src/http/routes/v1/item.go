package v1

import (
	itemController "belimang/src/http/controller/item"
	itemRepository "belimang/src/repositories/item"
	itemService "belimang/src/services/item"
)

func (i *V1Routes) MountItem() {
	gAdmin := i.Echo.Group("/admin")
	gMerchant := gAdmin.Group("/merchants")

	itemRepository := itemRepository.New(i.Db)
	itemService := itemService.New(itemRepository)
	itemController := itemController.New(itemService)

	gMerchant.POST("/:merchantId/items", itemController.CreateItem)
	gMerchant.GET("/:merchantId/items", itemController.GetAllItem)
}

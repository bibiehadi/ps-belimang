package v1

import (
	merchantController "belimang/src/http/controller/merchant"
	merchantRepository "belimang/src/repositories/merchant"
	merchantService "belimang/src/services/merchant"
)

func (i *V1Routes) MountMerchant() {
	gAdmin := i.Echo.Group("/admin")
	gMerchant := gAdmin.Group("/merchants")

	merchantRepository := merchantRepository.New(i.Db)
	merchantService := merchantService.New(merchantRepository)
	merchantController := merchantController.New(merchantService)

	gMerchant.GET("", merchantController.FindAll)
	gMerchant.POST("", merchantController.Create)
}

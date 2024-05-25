package v1

import (
	merchantController "belimang/src/http/controller/merchant"
	merchantRepository "belimang/src/repositories/merchant"
	merchantService "belimang/src/services/merchant"
)

func (i *V1Routes) MountMerchant() {
	gAdmin := i.Echo.Group("/admin")

	merchantRepository := merchantRepository.New(i.Db)
	merchantService := merchantService.New(merchantRepository)
	merchantController := merchantController.New(merchantService)

	gAdmin.POST("/merchants", merchantController.Create)
}

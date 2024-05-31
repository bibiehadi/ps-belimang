package v1

import (
	merchantController "belimang/src/http/controller/merchant"
	"belimang/src/http/middlewares"
	merchantRepository "belimang/src/repositories/merchant"
	merchantService "belimang/src/services/merchant"
)

func (i *V1Routes) MountMerchant() {
	gAdmin := i.Echo.Group("/admin")
	gAdminMerchant := gAdmin.Group("/merchants")
	gAdminMerchant.Use(middlewares.RequireAuth())
	merchantRepository := merchantRepository.New(i.Db)
	merchantService := merchantService.New(merchantRepository)
	merchantController := merchantController.New(merchantService)

	gAdminMerchant.GET("", merchantController.FindAll)
	gAdminMerchant.POST("", merchantController.Create)

	gMerchant := i.Echo.Group("/merchants")
	gMerchant.GET("/nearby/:latlong", merchantController.FindNearby)
}

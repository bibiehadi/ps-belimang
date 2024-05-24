package v1

import (
	authController "belimang/src/http/controller/auth"
	authRepository "belimang/src/repositories/auth"
	authService "belimang/src/services/auth"
)

func (i *V1Routes) MountAuth() {
	gAdmin := i.Echo.Group("/admin")
	gUser := i.Echo.Group("/user")

	authRepository := authRepository.New(i.Db)
	authService := authService.New(authRepository)
	authController := authController.New(authService)

	gAdmin.POST("/register", authController.RegisterAdmin)
	gAdmin.POST("/login", authController.Login)

	gUser.POST("/register", authController.RegisterUser)
	gUser.POST("/login", authController.Login)
}

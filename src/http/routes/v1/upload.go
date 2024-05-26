package v1

import (
	uploadController "belimang/src/http/controller/upload"
	"belimang/src/http/middlewares"
)

func (i *V1Routes) MountUpload() {
	g := i.Echo.Group("/image")
	g.Use(middlewares.RequireAuth())
	uploadController := uploadController.New()

	g.POST("", uploadController.UploadImage)
}

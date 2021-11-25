package routes

import (
	"keygo/app"
	"keygo/services/router"
)

func DefineApiRoute(router router.IRouter)  {
	keyController := app.ContainerKeyController

	router.POST("/set", keyController.Set)
	router.GET("/get", keyController.Get)
	router.DELETE("/del", keyController.Del)
	router.DELETE("/flush", keyController.Flush)
}

package routeslatihan

import (
	latihancontroller "latihan/app/controller"

	"github.com/gin-gonic/gin"
)

func Routes(ctx *gin.Engine) {
	routes := ctx.Group("latihan/")
	{
		// https://localhost:8000/latihan/list
		routes.GET("list", latihancontroller.Get)

		// https://localhost:8000/latihan/create
		routes.POST("create", latihancontroller.Post)

		// https://localhost:8000/latihan/update/:id
		routes.PATCH("update/:id", latihancontroller.Update)

		// https://localhost:8000/latihan/delete/:id
		routes.DELETE("delete/:id", latihancontroller.Delete)
	}
}

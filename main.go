package main

import (
	"latihan/app/routes/routeslatihan"

	"github.com/gin-gonic/gin"
)

func main() {
	Router := gin.Default()

	routeslatihan.Routes(Router)

	Router.Run(":8000")
}

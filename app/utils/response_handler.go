package utils

import "github.com/gin-gonic/gin"

type ResponCustom struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseHandler(c *gin.Context, statusServer int, statuscode bool, messagestring string, data interface{}) {
	var res ResponCustom

	res.Status = statuscode
	res.Message = messagestring
	res.Data = data

	c.JSON(statusServer, res)
}

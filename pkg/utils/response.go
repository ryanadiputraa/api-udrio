package utils

import (
	"github.com/gin-gonic/gin"
)

var status = map[int]string{
	200: "OK",
	201: "CREATED",
	400: "BAD_REQUEST",
	401: "UNAUTHORIZED",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	500: "INTERNAL_SERVER_ERROR",
}

func HttpResponse(code int, data interface{}) gin.H {
	return gin.H{
		"code":   code,
		"status": status[code],
		"error":  "",
		"data":   data,
	}
}

func HttpResponseError(code int, err string) gin.H {
	errMsg := map[string]string{
		"message": err,
	}

	return gin.H{
		"code":   code,
		"status": status[code],
		"error":  errMsg,
		"data":   nil,
	}
}

func HttpResponseWithMetaData(code int, data interface{}, meta interface{}) gin.H {
	return gin.H{
		"code":   code,
		"status": status[code],
		"error":  "",
		"data":   data,
		"meta":   meta,
	}
}

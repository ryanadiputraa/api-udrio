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

func HttpResponse(code int, errors interface{}, data interface{}) gin.H {
	if errors == nil {
		errors = ""
	}

	return gin.H{
		"code":   code,
		"status": status[code],
		"error":  errors,
		"data":   data,
	}
}

func HttpResponseWithMetaData(code int, errors interface{}, data interface{}, meta interface{}) gin.H {
	if errors == nil {
		errors = ""
	}

	return gin.H{
		"code":   code,
		"status": status[code],
		"error":  errors,
		"data":   data,
		"meta":   meta,
	}
}

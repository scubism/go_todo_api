package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
)

func AbortWithPublicError(c *gin.Context, code int, internalError error, publicErrorInfo interface{}) {

	var publicError error
	switch publicErrorInfo.(type) {
	case string:
		publicError = errors.New(publicErrorInfo.(string))
	default:
		publicError = internalError
	}

	log.Printf("ERROR: %s\n", publicError)
	if code >= 500 {
		// This the best way to log?
		trace := make([]byte, 1024)
		runtime.Stack(trace, true)
		log.Printf("%s", trace)
	}

	c.Error(internalError)
	c.AbortWithError(code, publicError).SetType(gin.ErrorTypePublic)
}

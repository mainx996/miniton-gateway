package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"www.miniton-gateway.com/pkg/log"
)

func Success(ctx *gin.Context, objs ...gin.H) {
	ctx.JSON(http.StatusOK, BuildSuccess(ctx, objs...))
}

func SuccessPure(ctx *gin.Context, objs ...gin.H) {
	ctx.PureJSON(http.StatusOK, BuildSuccess(ctx, objs...))
}

func BuildSuccess(ctx *gin.Context, objs ...gin.H) gin.H {
	var obj gin.H
	if len(objs) == 1 {
		obj = objs[0]
	} else {
		obj = make(gin.H)
	}

	if _, ok := obj["errno"]; !ok {
		obj["errno"] = 0
	}
	obj["traceid"] = log.TraceID(ctx.Request.Context())
	return obj
}

// Error http error
func Error(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"errno":   code,
		"errmsg":  msg,
		"traceid": log.TraceID(ctx.Request.Context()),
	})
}

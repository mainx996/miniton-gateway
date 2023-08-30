package http

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
	"www.miniton-gateway.com/pkg/log"
	schema2 "www.miniton-gateway.com/pkg/schema"
	"www.miniton-gateway.com/utils/ip"
	time2 "www.miniton-gateway.com/utils/time"
)

var DisableAccessLog = map[string]string{
	"/health": "",
}

func TraceID(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := log.NewTraceLog(c.Request.Context(), logger, c.GetHeader("X-Request-Id"))
		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set("X-Trace-Id", log.TraceID(ctx))
		c.Next()
	}
}

func AccessLog(c *gin.Context) {
	// 输入日志
	start := time2.GetCurrentTimestampByMill()
	body := ""
	if c.Request.Body != nil {
		rb, _ := c.GetRawData()
		if len(rb) > 0 {
			body = string(rb)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(rb))
		}
	}
	path := c.Request.URL.Path
	if _, ok := DisableAccessLog[path]; ok {
		c.Next()
		return
	}

	// 业务路由执行
	c.Next()

	// 输出日志
	latency := time2.GetCost(start)
	status := c.Writer.Status()
	method := c.Request.Method

	fields := []log.Field{
		zap.Int("status", status),
		zap.Int64("latency", latency),
		zap.String("method", method),
		zap.String("user-ip", ip.GetUserIP(c.Request.Header)),
		zap.String("query", c.Request.URL.RawQuery),
		zap.String("body", body),
	}
	// header携带token记录
	token := c.Request.Header.Get("token")
	if token != "" {
		fields = append(fields, zap.String("token", token))
	}

	// 响应打印
	if resp, exists := c.Get(schema2.RespKey); exists {
		if value, ok := resp.(schema2.RespStruct); ok {
			fields = append(fields, zap.Reflect("resp", value.Resp))
			fields = append(fields, zap.Reflect("errno", value.ErrorCode))
		} else {
			fields = append(fields, zap.Reflect("resp", resp))
		}
	}
	if len(c.Errors) > 0 {
		fields = append(fields, zap.Reflect("errs", c.Errors.Errors()))
	}
	log.Info(c.Request.Context(), path, fields...)
}

func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx := c.Request.Context()
			var brokenPipe bool
			if ne, ok := err.(*net.OpError); ok {
				if se, ok := ne.Err.(*os.SyscallError); ok {
					if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
						brokenPipe = true
					}
				}
			}
			httpRequest, _ := httputil.DumpRequest(c.Request, false)
			if brokenPipe {
				log.Error(ctx, c.Request.URL.Path,
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
				)
				c.Error(err.(error)) // nolint: errcheck
				c.Abort()
				return
			}
			log.Error(ctx, "[Recovery from panic]",
				zap.Time("time", time.Now()),
				zap.Any("error", err),
				zap.String("request", string(httpRequest)),
				zap.String("stack", string(debug.Stack())),
			)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	c.Next()
}

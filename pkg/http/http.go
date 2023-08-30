package http

import (
	"context"
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"www.miniton-gateway.com/pkg/config"
	"www.miniton-gateway.com/pkg/log"
)

var (
	Server *Serv
)

type (
	Serv struct {
		Router     *gin.Engine
		HttpServer http.Server
	}
)

func Init() {
	c := config.Config.HTTPConfig
	e := gin.New()
	if config.Mode == config.ProdMode {
		gin.SetMode(gin.ReleaseMode)
	}
	e.Use(TraceID(log.Log))
	e.Use(AccessLog)
	e.Use(Recovery)
	pprof.RouteRegister(&(e.RouterGroup), "ton/debug/pprof")
	e.GET("/health", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	addr := fmt.Sprintf("0.0.0.0:%d", c.Port)
	Server = &Serv{
		Router: e,
	}
	Server.HttpServer = http.Server{Addr: addr, Handler: Server.Router}
}

func Start() {
	log.Info(context.Background(), "http server starting ...")
	go func() {
		if err := gracehttp.Serve(&Server.HttpServer); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func Stop() {
	c := config.Config.HTTPConfig
	log.Info(context.Background(), "http server stopping ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.ShutdownTimeout))
	defer cancel()
	_ = Server.HttpServer.Shutdown(ctx)
}

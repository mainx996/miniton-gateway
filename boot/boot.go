package boot

import (
	"context"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"www.miniton-gateway.com/app"
	"www.miniton-gateway.com/pkg/config"
	"www.miniton-gateway.com/pkg/http"
	"www.miniton-gateway.com/pkg/log"
	"www.miniton-gateway.com/pkg/mysql"
)

func Init(mode string) {
	config.Init(mode)
	log.Init()
	mysql.Init()
	http.Init()
	app.Init()
}

func Run() {
	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		http.Start()
		return nil
	})
	_ = g.Wait()
}

func AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
	http.Stop()
	log.Info(context.Background(), "stop done")
}

package app

import (
	"www.miniton-gateway.com/app/controller/auth"
	"www.miniton-gateway.com/app/controller/demo"
	"www.miniton-gateway.com/pkg/http"
)

func Init() {
	demoRoute()
}

func demoRoute() {
	r := http.Server.Router
	r.Use(auth.Auth())
	d := r.Group("/demo")
	d.GET("/detail", demo.Detail)
	d.GET("/list", demo.List)
	d.POST("/create", demo.Create)
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gavinzhou/hello-gin/pkg/setting"

	"github.com/gavinzhou/hello-gin/routers"
)

func main() {
	config := setting.Config.ServerConfig
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.HTTPPort),
		Handler:        router,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

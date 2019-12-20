package main

import (
	"context"
	"apiservicedemo/config"
	"apiservicedemo/handlers"
	"apiservicedemo/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	serverClosed chan interface{}
)

func main() {
	appConf := config.AppConfig("./config.json")
	serverClosed = make(chan interface{})

	logger.Init(appConf.Mode)
	gin.SetMode(appConf.Mode)

	logger.Info("Server Starting")

	// Register Handler
	router := gin.Default()
	router.Use(gin.Recovery())

	{
		router.GET("health_check", handlers.HealthCheck)
	}

	srv := &http.Server{
		Addr:    ":" + appConf.Port,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdowning Server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go func() {
		defer close(serverClosed)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Fatal("Server Shutdown failed:" + err.Error())
		}
	}()

	select {
	case <-serverClosed:
		logger.Info("Shutdown Server ok")
	case <-ctx.Done():
		logger.Warn("timeout of 5 seconds.")
	}
	logger.Info("ByeBye")
}

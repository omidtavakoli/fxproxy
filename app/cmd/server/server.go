package main

import (
	"app/internal/http/rest"
	"app/internal/logger"
	"app/internal/proxy"
	"context"
	"github.com/go-playground/validator/v10"
	"os"
	"sync"
)

type Server struct {
	sync.WaitGroup
	Config      *MainConfig
	RESTHandler *rest.Handler
	Logger      *logger.StandardLogger
}

func NewServer(cfg *MainConfig, logger *logger.StandardLogger) *Server {
	return &Server{
		Config: cfg,
		Logger: logger,
	}
}

// Initialize is responsible for app initialization and wrapping required dependencies
func (s *Server) Initialize(ctx context.Context) error {
	v := validator.New()
	service := proxy.CreateService(&s.Config.Proxy, s.Logger, v)
	handler := rest.CreateHandler(service)
	s.RESTHandler = handler
	return nil
}

// Start starts the application in blocking mode
func (s *Server) Start(ctx context.Context) {
	// Start TimerBased CRON Jobs
	//go app.StartCronJobs(ctx)

	// Create Router for HTTP Server
	router := SetupRouter(s.RESTHandler, s.Config)
	//s.RESTHandler.Prometheus = prometheus
	//logrus.Info(prometheus)

	// Start REST Server in Blocking mode
	s.RESTHandler.Start(ctx, s.Config.Server.Port, router)
}

// GracefulShutdown listen over the quitSignal to graceful shutdown the app
func (s *Server) GracefulShutdown(quitSignal <-chan os.Signal, done chan<- bool) {
	// const op = "app.gacefulshutdown"
	// Wait for OS signals
	<-quitSignal

	// Kill the API Endpoints first
	s.RESTHandler.Stop()

	close(done)
}

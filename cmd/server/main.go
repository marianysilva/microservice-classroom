package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sumelms/microservice-classroom/internal/classroom"
	"github.com/sumelms/microservice-classroom/internal/shared"
	"github.com/sumelms/microservice-classroom/internal/shared/clients"
	courses "github.com/sumelms/microservice-classroom/internal/shared/clients/rpc"
	"github.com/sumelms/microservice-classroom/pkg/config"
	database "github.com/sumelms/microservice-classroom/pkg/database/postgres"
	log "github.com/sumelms/microservice-classroom/pkg/logger"
	"github.com/sumelms/microservice-classroom/swagger"
	"golang.org/x/sync/errgroup"
)

var (
	logger     log.Logger
	httpServer *shared.Server
)

//nolint:funlen
func main() {
	// Logger
	logger = log.NewLogger()
	logger.Log("msg", "service started")

	// Configuration
	cfg, err := loadConfig()
	if err != nil {
		logger.Log("msg", "exit", "error", err)
		os.Exit(-1)
	}

	// Database
	db, err := database.Connect(cfg.Database)
	if err != nil {
		logger.Log("msg", "database error", "error", err)
		os.Exit(1)
	}

	// Initialize the domain services
	svcLogger := logger.With("component", "service")

	// Initializing the HTTP Services
	rpcLogger := logger.With("component", "rpc")

	coursesClient, err := clients.NewCoursesClient(cfg, rpcLogger)
	if err != nil {
		logger.Log("msg", "unable to connect with course service", "error", err)
		os.Exit(1)
	}
	coursesClientService := courses.NewCoursesService(coursesClient)
	logger.Log("msg", "connection with courses client service started")

	classroomSvc, err := classroom.NewService(db, svcLogger.Logger(), clients.ClientServices{Courses: *coursesClientService})
	if err != nil {
		logger.Log("msg", "unable to start classroom service", "error", err)
		os.Exit(1)
	}

	// Gracefully shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		// Initialize the router
		router := mux.NewRouter().StrictSlash(true)
		// Global Middlewares
		router.Use(shared.CorsMiddleware)

		// Register Swagger handler
		swagger.Register(router)

		// Initializing the HTTP Services
		httpLogger := logger.With("component", "http")

		if err = classroom.NewHTTPService(router, classroomSvc, httpLogger.Logger()); err != nil {
			logger.Log("msg", "unable to start a service: classroom", "error", err)
			return err
		}

		// Create the HTTP Server
		httpServer, err = shared.NewServer(cfg.Server.HTTP, router, httpLogger)
		if err != nil {
			return err
		}

		return httpServer.Start()
	})

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	logger.Log("msg", "received shutdown signal")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if httpServer != nil {
		httpServer.Stop(shutdownCtx)
	}

	if err := g.Wait(); err != nil {
		logger.Log("msg", "server returning an error", "error", err)
		defer os.Exit(2)
	}

	logger.Log("msg", "service ended")
}

func loadConfig() (*config.Config, error) {
	// Configuration
	configPath := os.Getenv("SUMELMS_CONFIG_PATH")
	if configPath == "" {
		configPath = "./config.yml"
	}

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

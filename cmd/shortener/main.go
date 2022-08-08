package main

import (
	"context"
	"github.com/ChristinaFomenko/shortener/configs"
	"github.com/ChristinaFomenko/shortener/internal/app/generator"
	"github.com/ChristinaFomenko/shortener/internal/app/hasher"
	repositoryURL "github.com/ChristinaFomenko/shortener/internal/app/repository/urls"
	authService "github.com/ChristinaFomenko/shortener/internal/app/service/auth"
	pingService "github.com/ChristinaFomenko/shortener/internal/app/service/ping"
	serviceURL "github.com/ChristinaFomenko/shortener/internal/app/service/urls"
	"github.com/ChristinaFomenko/shortener/internal/app/worker"
	"github.com/ChristinaFomenko/shortener/internal/handlers"
	"github.com/ChristinaFomenko/shortener/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// Config
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("failed to retrieve env variables, %v", err)
	}

	wp := worker.New(ctx, cfg.NumbWorkers, cfg.WorkerBuff)
	go func() {
		wp.Run(ctx)
	}()

	// Repositories
	repository, err := repositoryURL.NewStorage(cfg.FileStoragePath, cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("failed to create a storage %v", err)
	}
	//defer func(repository repositoryURL.Repo) {
	//	_ = repository.Close()
	//}(repository)

	// Services
	helper := generator.NewGenerator()
	hash := hasher.NewHasher(cfg.SecretKey)
	service := serviceURL.NewService(repository, helper, cfg.BaseURL)
	authSrvc := authService.NewService(helper, hash)
	pingSrvc := pingService.NewService(repository)

	// Route
	router := chi.NewRouter()

	compress, err := middlewares.NewCompressor()
	if err != nil {
		log.Fatalf("compressor failed %v", err)
	}

	auth := middlewares.NewAuthenticator(authSrvc)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middlewares.Decompressing)
	router.Use(compress.Compressing)
	router.Use(auth.Auth)

	router.Post("/", handlers.New(service, auth, pingSrvc, wp).Shorten)
	router.Get("/{id}", handlers.New(service, auth, pingSrvc, wp).Expand)
	router.Post("/api/shorten", handlers.New(service, auth, pingSrvc, wp).APIJSONShorten)
	router.Get("/api/user/urls", handlers.New(service, auth, pingSrvc, wp).FetchURLs)
	router.Get("/ping", handlers.New(service, auth, pingSrvc, wp).Ping)
	router.Post("/api/shorten/batch", handlers.New(service, auth, pingSrvc, wp).ShortenBatch)
	router.Delete("/api/user/urls", handlers.New(service, auth, pingSrvc, wp).DeleteUserURLs)

	address := cfg.ServerAddress
	log.WithField("address", address).Info("server starts")

	log.Fatal(http.ListenAndServe(address, router))

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	select {
	case <-sigint:
		cancel()
	case <-ctx.Done():
	}
}

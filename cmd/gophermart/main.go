package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DEHbNO4b/practicum_project/internal/handlers"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/DEHbNO4b/practicum_project/internal/repository/postgres"
	"github.com/DEHbNO4b/practicum_project/internal/services"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

func main() {
	if err := logger.Initialize("info"); err != nil {
		panic(err)
	}
	if err := run(); err != nil {
		logger.Log.Fatal(err.Error())
		os.Exit(0)
	}
}
func run() error {
	cfg := parseFlag()
	pdb, err := postgres.NewUserDB(cfg.Database_url)
	if err != nil {
		return err
	}
	userService := services.NewUserService(pdb)
	uhandler := handlers.NewRegister(&userService)
	r := chi.NewRouter()
	r.Post(`/api/user/register`, uhandler.Register)
	srv := &http.Server{
		Addr:    cfg.Run_adress,
		Handler: r,
	}
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()
	logger.Log.Info("Running server", zap.String("adress", cfg.Run_adress))
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Log.Fatal("HTTP server ListenAndServe Error", zap.Error(err))
	}
	<-stopped
	logger.Log.Info("Have a nice day!")
	return nil
}

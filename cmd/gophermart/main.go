package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DEHbNO4b/practicum_project/internal/agent"
	"github.com/DEHbNO4b/practicum_project/internal/config"
	"github.com/DEHbNO4b/practicum_project/internal/handlers/balance"
	"github.com/DEHbNO4b/practicum_project/internal/handlers/debit"
	"github.com/DEHbNO4b/practicum_project/internal/handlers/order"
	"github.com/DEHbNO4b/practicum_project/internal/handlers/user"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/DEHbNO4b/practicum_project/internal/middleware/authentication"
	"github.com/DEHbNO4b/practicum_project/internal/service"
	"github.com/DEHbNO4b/practicum_project/internal/storage"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctx := context.Background()

	//config
	cfg := config.Get()

	//logger
	if err := logger.Initialize("info"); err != nil {
		return err
	}

	//init repository store(with postgres inside)
	storage, err := storage.New(ctx)
	if err != nil {
		return err
	}
	defer storage.Close() //close DBs

	//init service manager(with storage inside)
	serviceManager, err := service.NewManager(ctx, storage)
	if err != nil {
		return fmt.Errorf("%s %w", "unable to create service manager", err)
	}
	//init and start AccrualAgentManager
	agentManager := agent.NewManager(ctx, storage)
	go agentManager.Start()

	//create handlers(with serviceManager inside)
	uhandler := user.NewUsers(ctx, serviceManager)
	oHandler := order.NewOrders(ctx, serviceManager)
	bHandler := balance.NewBalance(ctx, serviceManager)
	dHandler := debit.NewDebit(ctx, serviceManager)

	//init router(chi)
	router := chi.NewRouter()
	router.Post(`/api/user/register`, uhandler.Register)
	router.Post(`/api/user/login`, uhandler.Login)
	router.Route(`/api/user`, func(r chi.Router) {
		r.Use(authentication.Auth)
		r.Post("/orders", oHandler.LoadOrder)
		r.Get("/orders", oHandler.GetOrders)
		r.Get("/balance", bHandler.GetBalance)
		r.Post("/balance/withdraw", dHandler.AddDebit)
		r.Get("/withdrawals", dHandler.GetAllDebits)
	})
	srv := &http.Server{
		Addr:         cfg.RunAdress,
		Handler:      router,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}

	//gracefull shutdown
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

	//running server
	logger.Log.Info("Running server", zap.String("adress", cfg.RunAdress))
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Log.Fatal("HTTP server ListenAndServe Error", zap.Error(err))
	}

	<-stopped
	logger.Log.Info("Have a nice day!")

	return nil
}

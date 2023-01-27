package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/FirdavsMF/wallet-api/config"
	"github.com/FirdavsMF/wallet-api/internal/usecase"
	"github.com/FirdavsMF/wallet-api/pkg/logger"
)

func main() {

	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	conn, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.PG.HOST, config.PG.Port, config.PG.Username, config.PG.Password, config.DbName))
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		Username: config.Redis.Username,
		Password: config.Redis.Password,
	})
	l := logger.New(config.Log.Level)
	u := usecase.NewUsecase(conn, l, redisClient)
	webAPIHandler := usecase.NewWebApiHandler(u)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.New()
	r.POST("/login", webAPIHandler.Login)
	v1 := r.Group("/v1")
	{
		v1.Use(webAPIHandler.HeaderCheck())
		v1.POST("/checkwallet", webAPIHandler.CheckWallet)
		v1.POST("/charge", webAPIHandler.Charge)
		v1.POST("/gethistory", webAPIHandler.GetWalletHistory)
		v1.POST("/getbalance", webAPIHandler.GetWalletBalance)
	}

	srv := &http.Server{
		Addr:    config.HTTP.Port,
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

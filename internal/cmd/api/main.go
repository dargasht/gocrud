package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bytedance/sonic"
	"github.com/dargasht/gocrud"
	"github.com/dargasht/gocrud/internal/cfg"
	"github.com/dargasht/gocrud/internal/database/repo"
	"github.com/dargasht/gocrud/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "go.uber.org/automaxprocs" // use this it is good
	"go.uber.org/zap"
)

func main() {

	//---------------------------------------------------------------
	//Setup logger

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//---------------------------------------------------------------
	//Setup db

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.CON_STRING)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.String("error", err.Error()))
		os.Exit(1)
	}
	defer pool.Close()
	db := repo.New(pool)
	s, _ := pool.Acquire(ctx)
	defer s.Release()

	//---------------------------------------------------------------
	//Setup sqlx db if needed

	// _, err = store.NewPostgresDB()
	// if err != nil {
	// 	logger.Fatal("Failed to connect to database", zap.String("error", err.Error()))
	// 	os.Exit(1)
	// }

	//---------------------------------------------------------------
	// Setup app and router
	app := fiber.New(fiber.Config{
		ErrorHandler: gocrud.CRUDErrorHandler(logger),
		AppName:      "Some CRUD App",
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
	})

	router.SetupApp(app, db, logger)

	//---------------------------------------------------------------
	// cron jobs if any

	//---------------------------------------------------------------

	go func() {
		if err = app.Listen(":" + cfg.PORT); err != nil {
			logger.Fatal("Failed to start server", zap.String("error", err.Error()))
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	logger.Info("Gracefully shutting down...")
	err = app.Shutdown()
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("Running cleanup tasks...")

	// Your cleanup tasks go here

	logger.Info("Fiber was successful shutdown.")
}

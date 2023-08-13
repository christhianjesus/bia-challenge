package main

import (
	"database/sql"
	"fmt"

	"github.com/christhianjesus/bia-challenge/cmd/consumption/config"
	consumptionApp "github.com/christhianjesus/bia-challenge/internal/application/consumption"
	"github.com/christhianjesus/bia-challenge/internal/application/period/strategies"
	"github.com/christhianjesus/bia-challenge/internal/infrastructure/consumption"
	"github.com/joeshaw/envdecode"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	conf := &config.Config{}

	if err := envdecode.Decode(conf); err != nil {
		panic(fmt.Errorf("Cannot read from env: %w", err))
	}

	// db
	db, err := sql.Open("postgres", conf.DSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	setupHandlers(conf, e.Group("/api"), db)

	// Start server
	e.Logger.Fatal(e.Start(conf.Addr()))
}

func setupHandlers(conf *config.Config, router *echo.Group, db *sql.DB) {

	psf := &strategies.PeriodStrategyFactory{}

	// Repos
	consumptionRepository := consumption.NewPostgreSQLConsumptionRepository(db)

	// Services
	consumptionService := consumptionApp.NewConsumptionService(consumptionRepository, nil, psf)

	// Handlers
	consumptionHandler := consumption.NewConsumptionHandler(consumptionService)

	// Register routes
	consumptionHandler.RegisterRoutes(router)
}

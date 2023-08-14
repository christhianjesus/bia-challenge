package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/christhianjesus/bia-challenge/cmd/consumption/config"
	addressApp "github.com/christhianjesus/bia-challenge/internal/application/address"
	consumptionApp "github.com/christhianjesus/bia-challenge/internal/application/consumption"
	"github.com/christhianjesus/bia-challenge/internal/infrastructure/address"
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

	// Repos
	addressRepository := address.NewMSAddressRepository(http.DefaultClient)
	consumptionRepository := consumption.NewPostgreSQLConsumptionRepository(db)

	// Services
	addressService := addressApp.NewAddressService(addressRepository)
	consumptionService := consumptionApp.NewConsumptionService(consumptionRepository)
	consumptionPeriodsService := consumptionApp.NewConsumptionPeriodsService()

	// Handlers
	consumptionHandler := consumption.NewConsumptionHandler(consumptionService, consumptionPeriodsService, addressService)

	// Register routes
	consumptionHandler.RegisterRoutes(router)
}

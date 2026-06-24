package main

import (
	"backend/db"
	"backend/internal/config"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/router"
	"backend/internal/service"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	ctx := context.Background()

	if err != nil {
		log.Fatal("not able to load .env file")
	}

	pgxPool, err := config.InitDatabase(ctx, cfg)
	if err != nil {
		fmt.Errorf("error :%v", err)
	}

	defer pgxPool.Close()

	queries := db.New(pgxPool)

	driverRepo := repository.NewDriverRepository(queries)

	driverService := service.NewDriverService(driverRepo)

	driverHandler := handler.NewDriverHandler(driverService)

	r := router.SetupRouter(driverHandler)
	log.Println("Server running on port :8000")
	http.ListenAndServe(":8000", r)
}

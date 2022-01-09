package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/haton14/ohagi-api/controller"
	"github.com/haton14/ohagi-api/ent"
	"github.com/haton14/ohagi-api/infrastructure/datastore"
	"github.com/haton14/ohagi-api/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Controller struct {
	dbClient *ent.Client
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	//DataStore
	db := datastore.NewDB(os.Getenv("DATABASE_URL"))
	dbClient, err := db.Open()
	if err != nil {
		log.Fatalf("database open err. %s", err)
	}
	defer dbClient.Close()

	//Migration
	if err := dbClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Usecase
	createRecord := usecase.NewCreateRecord(dbClient)

	// Controller
	recordController := controller.NewRecord(dbClient, createRecord)

	// Routes
	e.GET("/records", recordController.List)
	e.POST("records", recordController.Create)

	// Set port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Start server
	go func() {
		if err := e.Start(":" + port); err != nil {
			e.Logger.Error("shutting down the server:", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("chan")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

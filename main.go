package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Records struct {
		Records []Recode `json:"records"`
	}
	Recode struct {
		DateTime string `json:"date_time"`
		Foods    []Food `json:"foods"`
	}
	Food struct {
		FoodName string `json:"food_name"`
		Weight   int    `json:"weight"`
	}
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/records/today", hello)
	e.GET("/records/recent", hello)
	e.GET("/records/year/:year/month/:month", hello)

	// Start server
	go func() {
		if err := e.Start(":8000"); err != nil {
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

// Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

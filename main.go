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
		Years []YearRecord `json:"years"`
	}

	YearRecord struct {
		Year   int           `json:"year"`
		Months []MonthRecord `json:"months"`
	}

	MonthRecord struct {
		Month int         `json:"month"`
		Days  []DayRecord `json:"days"`
	}

	DayRecord struct {
		Day    int    `json:"day"`
		Record Record `json:"record"`
	}

	Record struct {
		ID            int    `json:"id"`
		Foods         []Food `json:"foods"`
		LastUpdatedAt int64  `json:"last_updated_at"`
		CreatedAt     int64  `json:"created_at"`
	}

	Food struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Amount        int    `json:"amount"`
		Unit          string `json:"unit"`
		LastUpdatedAt int64  `json:"last_updated_at,omitempty"`
	}
)

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

	// Routes
	e.GET("/records/today", hello)
	e.GET("/records/recent", hello)
	e.GET("/records/year/:year/month/:month", hello)
	e.GET("/records", topRecords)
	e.POST("records", postRecord)

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

// Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func topRecords(c echo.Context) error {
	records := &Records{
		Years: []YearRecord{
			{
				Year: 2021,
				Months: []MonthRecord{
					{
						Month: 1,
						Days: []DayRecord{{
							Day: 1,
							Record: Record{
								ID: 123,
								Foods: []Food{{
									ID:     12345,
									Name:   "ペレット",
									Amount: 10,
									Unit:   "g",
								}},
								LastUpdatedAt: 1234567890,
								CreatedAt:     1234567890,
							},
						}},
					},
				},
			},
		},
	}
	return c.JSON(http.StatusOK, records)
}

func postRecord(c echo.Context) error {
	record := Record{}
	if err := c.Bind(record); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, record)
}

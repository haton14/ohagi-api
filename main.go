package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/haton14/ohagi-api/infrastructure/datastore"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Records struct {
		Records []Record `json:"records"`
	}

	Record struct {
		ID            int    `json:"id,omitempty"`
		Foods         []Food `json:"foods"`
		LastUpdatedAt int64  `json:"last_updated_at,omitempty"`
		CreatedAt     int64  `json:"created_at,omitempty"`
	}

	Food struct {
		ID            int    `json:"id,omitempty"`
		Name          string `json:"name"`
		Amount        int    `json:"amount,omitempty"`
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

	// Routes
	e.GET("/records", getRecords)
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

func getRecords(c echo.Context) error {
	records := Records{
		[]Record{
			{

				ID: 121,
				Foods: []Food{{
					ID:     12344,
					Name:   "ペレット",
					Amount: 10,
					Unit:   "g",
				}},
				LastUpdatedAt: 1638280364,
				CreatedAt:     1638280364,
			},
			{

				ID: 122,
				Foods: []Food{{
					ID:     12344,
					Name:   "ペレット",
					Amount: 10,
					Unit:   "g",
				}},
				LastUpdatedAt: 1638366764,
				CreatedAt:     1638366764,
			},
			{

				ID: 123,
				Foods: []Food{{
					ID:     12344,
					Name:   "ペレット",
					Amount: 10,
					Unit:   "g",
				}},
				LastUpdatedAt: 1640181164,
				CreatedAt:     1640181164,
			},
			{

				ID: 124,
				Foods: []Food{{
					ID:     12344,
					Name:   "ペレット",
					Amount: 10,
					Unit:   "g",
				}},
				LastUpdatedAt: 1640094764,
				CreatedAt:     1640094764,
			},
			{

				ID: 125,
				Foods: []Food{{
					ID:     12345,
					Name:   "ペレット",
					Amount: 10,
					Unit:   "g",
				}},
				LastUpdatedAt: 1640958764,
				CreatedAt:     1640958764,
			},
			{

				ID: 126,
				Foods: []Food{{
					ID:     12345,
					Name:   "ペレット",
					Amount: 10,
					Unit:   "g",
				}},
				LastUpdatedAt: 1641045164,
				CreatedAt:     1641045164,
			},
		},
	}
	return c.JSON(http.StatusOK, &records)

}

func postRecord(c echo.Context) error {
	record := Record{}
	if err := c.Bind(&record); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, record)
}

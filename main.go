package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/haton14/ohagi-api/ent"
	"github.com/haton14/ohagi-api/ent/recordfood"
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
		ID            *int   `json:"id,omitempty"`
		Name          string `json:"name"`
		Amount        int    `json:"amount,omitempty"`
		Unit          string `json:"unit"`
		LastUpdatedAt *int64 `json:"last_updated_at,omitempty"`
	}
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
	controller := Controller{dbClient: dbClient}

	//Migration
	if err := dbClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Routes
	e.GET("/records", controller.getRecords)
	e.POST("records", controller.postRecord)

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

func (controller *Controller) getRecords(c echo.Context) error {
	rq := controller.dbClient.Record.Query().Limit(50)
	recordsEnt, err := rq.All(context.Background())
	if err != nil {
		c.Logger().Error("All: ", err)
		return c.String(http.StatusInternalServerError, "All: "+err.Error())
	}
	records := []Record{}
	ids := []int{}
	for _, r := range recordsEnt {
		records = append(records, Record{
			ID:            r.ID,
			Foods:         []Food{},
			CreatedAt:     r.CreatedAt.Unix(),
			LastUpdatedAt: r.LastUpdatedAt.Unix(),
		})
		ids = append(ids, r.ID)
	}
	recordFoodsEnt, err := controller.dbClient.RecordFood.Query().
		Where(func(s *sql.Selector) { sql.InInts(recordfood.FieldRecordID, ids...) }).
		All(context.Background())
	if err != nil {
		c.Logger().Error("All: ", err)
		return c.String(http.StatusInternalServerError, "All: "+err.Error())
	}
	foodsEnt, err := controller.dbClient.Food.Query().All(context.Background())
	if err != nil {
		c.Logger().Error("All: ", err)
		return c.String(http.StatusInternalServerError, "All: "+err.Error())
	}

	for i, r := range records {
		foods := []Food{}
		for _, rf := range recordFoodsEnt {
			if r.ID == rf.RecordID {
				for _, f := range foodsEnt {
					if rf.FoodID == f.ID {
						food := Food{
							Name:   f.Name,
							Amount: rf.Amount,
							Unit:   f.Unit,
						}
						foods = append(foods, food)
						continue
					}
				}
			}
		}
		r.Foods = append(r.Foods, foods...)
		records[i] = r
	}

	return c.JSON(http.StatusOK, &Records{records})

}

func (controller *Controller) postRecord(c echo.Context) error {
	record := Record{}
	if err := c.Bind(&record); err != nil {
		return err
	}
	recordEnt, err := controller.dbClient.Record.Create().SetCreatedAt(time.Now()).SetLastUpdatedAt(time.Now()).Save(context.Background())
	if err != nil {
		c.Logger().Error("Save: ", err)
		return c.String(http.StatusInternalServerError, "Save: "+err.Error())
	}
	foodsEnt, err := controller.dbClient.Food.Query().All(context.Background())
	if err != nil {
		c.Logger().Error("All: ", err)
		return c.String(http.StatusInternalServerError, "All: "+err.Error())
	}
	recordFoodBulk := make([]*ent.RecordFoodCreate, len(record.Foods))
	for i, food := range record.Foods {
		match := false
		for _, foodEnt := range foodsEnt {
			if food.Name == foodEnt.Name && food.Unit == foodEnt.Unit {
				recordFoodBulk[i] = controller.dbClient.RecordFood.Create().SetRecordID(recordEnt.ID).SetFoodID(foodEnt.ID).SetAmount(food.Amount)
				match = true
				break
			}
		}
		if !match {
			foodEnt, err := controller.dbClient.Food.Create().SetName(food.Name).SetUnit(food.Unit).Save(context.Background())
			if err != nil {
				c.Logger().Error("All: ", err)
				return c.String(http.StatusInternalServerError, "All: "+err.Error())
			}
			recordFoodBulk[i] = controller.dbClient.RecordFood.Create().SetRecordID(recordEnt.ID).SetFoodID(foodEnt.ID).SetAmount(food.Amount)
		}
	}
	_, err = controller.dbClient.RecordFood.CreateBulk(recordFoodBulk...).Save(context.Background())
	if err != nil {
		c.Logger().Error("All: ", err)
		return c.String(http.StatusInternalServerError, "All: "+err.Error())
	}

	record.ID = recordEnt.ID
	record.CreatedAt = recordEnt.CreatedAt.Unix()
	record.LastUpdatedAt = recordEnt.LastUpdatedAt.Unix()

	return c.JSON(http.StatusCreated, record)
}

package datastore

import (
	"fmt"
	"strconv"

	"github.com/haton14/ohagi-api/ent"
	_ "github.com/lib/pq"
)

type DB struct {
	host     string
	port     int
	user     string
	dbname   string
	password string
}

func NewDB(host string, port string, user, dbname, password string) DB {
	numPort, _ := strconv.Atoi(port)
	return DB{host, numPort, user, dbname, password}
}
func (d *DB) Open() (*ent.Client, error) {
	return ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		d.host, d.port, d.user, d.dbname, d.password))
}

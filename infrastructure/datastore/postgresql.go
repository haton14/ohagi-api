package datastore

import (
	"github.com/haton14/ohagi-api/ent"
	_ "github.com/lib/pq"
)

type DB struct {
	databaseURL string
}

func NewDB(databaseURL string) *DB {
	return &DB{databaseURL}
}
func (d *DB) Open() (*ent.Client, error) {
	return ent.Open("postgres", d.databaseURL)
}

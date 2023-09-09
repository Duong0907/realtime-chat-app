// FRAMEWORKS AND DRIVERS LAYER
// Software areas that reside inside this layer are:
// - User Interface
// - Database
// - External Interfaces (eg: Native platform API)
// - Web (eg: Network Request)
// - Devices (eg: Printers and Scanners)

package db

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)


type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	connStr := "postgres://root:secret@localhost:5433/root?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	log.Println("Connect to database successfully")
	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

// sql.DB need to implement DBTX interface
// But this is default
func (d *Database) GetDB() *sql.DB {
	return d.db
}
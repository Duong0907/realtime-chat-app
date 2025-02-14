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
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)


type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	// Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Get the connection string from the environment variable
    connStr := os.Getenv("DB_CONN_STR")
    if connStr == "" {
        log.Fatal("DB_CONN_STR environment variable not set")
    }

	
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
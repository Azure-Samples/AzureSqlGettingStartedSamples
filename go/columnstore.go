package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/microsoft/go-mssqldb/azuread"
)

// Replace with your own connection parameters
var server = "your_server.database.windows.net"
var database = "your_database"

var db *sql.DB

// Delete an employee from database
func ExecuteAggregateStatement(db *sql.DB) {
	ctx := context.Background()

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	var result string

	// Execute long non-query to aggregate rows
	err = db.QueryRowContext(ctx, "SELECT SUM(Price) as sum FROM Table_with_3M_rows").Scan(&result)
	if err != nil {
		log.Fatal("Error executing query: " + err.Error())
	}

	fmt.Printf("Sum: %s\n", result)
}

func main() {
	// Connect to database
	connString := fmt.Sprintf("server=%s;database=%s;fedauth=ActiveDirectoryDefault;", server, database)
	var err error

	// Create connection pool
	db, err = sql.Open(azuread.DriverName, connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Printf("Connected!\n")

	defer db.Close()

	t1 := time.Now()
	fmt.Printf("Start time: %s\n", t1)

	ExecuteAggregateStatement(db)

	t2 := time.Since(t1)
	fmt.Printf("The query took: %s\n", t2)
}

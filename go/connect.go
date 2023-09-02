package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/microsoft/go-mssqldb/azuread"
)

// Replace with your own connection parameters
var server = "your_server.database.windows.net"
var database = "your_database"

var db *sql.DB

func main() {
	var err error

        // Build connection string
        connString := fmt.Sprintf("server=%s;database=%s;fedauth=ActiveDirectoryDefault;", server, database)

        // Create connection pool
        db, err = sql.Open(azuread.DriverName, connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")

	// Close the database connection pool after program executes
	defer db.Close()

	SelectVersion()
}

// Gets and prints SQL Server version
func SelectVersion() {
	// Use background context
	ctx := context.Background()

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	stmt, err := db.Prepare("select @@version")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var result string
	err = row.Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("%s\n", result)
}

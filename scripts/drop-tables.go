package scripts

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func main() {
	dbURI := os.Getenv("DB_URI")
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}
	defer db.Close()

	// Get all table names
	rows, err := db.Query("SELECT tablename FROM pg_tables WHERE schemaname='public'")
	if err != nil {
		log.Fatalf("Failed to query table names: %v\n", err)
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			log.Fatalf("Failed to scan table name: %v\n", err)
		}
		tables = append(tables, table)
	}

	// Drop all tables
	for _, table := range tables {
		if _, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", table)); err != nil {
			log.Fatalf("Failed to drop table %s: %v\n", table, err)
		}
		fmt.Printf("Dropped table %s\n", table)
	}
}

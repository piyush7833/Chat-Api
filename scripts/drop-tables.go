package scripts

import (
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"github.com/piyush7833/Chat-Api/services"
)

func DroptTables() {
	tables := []string{"User", "FriendRequest", "Friend", "Media", "Message", "Notification", "Reminder", "Status", "Call", "Block", "Tag"}

	// Drop all tables
	for _, table := range tables {
		query := fmt.Sprintf(`DROP TABLE IF EXISTS "%s" CASCADE`, table)
		if _, err := services.Db.Exec(query); err != nil {
			log.Fatalf("Failed to drop table %s: %v\n", table, err)
		}
		// fmt.Printf("Dropped table %s\n", table)
	}
}

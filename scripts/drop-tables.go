package scripts

import (
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"github.com/piyush7833/Chat-Api/services"
)

func DroptTables() {
	tables := []string{"users", "userRelation", "messages", "notificationId", "reminders", "statuses", "calls", "tags", "visibility", "groups", "groupUsers", "messageThreads"}
	types := []string{"status_type", "call_type", "group_user_role", "message_type", "relation_status", "visibility_type"}

	// Drop all tables

	for _, table := range tables {
		query := fmt.Sprintf(`DROP TABLE IF EXISTS "%s" CASCADE`, table)
		if _, err := services.Db.Exec(query); err != nil {
			log.Fatalf("Failed to drop table %s: %v\n", table, err)
		}
		// fmt.Printf("Dropped table %s\n", table)
	}

	for _, types := range types {
		query := fmt.Sprintf(`DROP TYPE IF EXISTS "%s"`, types)
		if _, err := services.Db.Exec(query); err != nil {
			log.Fatalf("Failed to drop type %s: %v\n", types, err)
		}
	}
}

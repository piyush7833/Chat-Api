package scripts

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/piyush7833/Chat-Api/services"
)

func CreateTables() {
	// Get a list of all SQL files in the current directory
	var dir string
	var err error
	if os.Getenv("ENVIRONMENT") == "test" || os.Getenv("ENVIRONMENT") == "github_test" {
		childDir, err := os.Getwd()
		dir = filepath.Dir(childDir)
		if err != nil {
			log.Fatal("Error getting current working directory:", err)
		}
	} else {
		dir, err = os.Getwd()
		if err != nil {
			log.Fatal("Error getting current working directory:", err)
		}
	}

	modelsDir := filepath.Join(dir, "models")
	if err != nil {
		log.Fatal("Error reading directory:", err)
	}
	files := []string{"users.sql", "user-relation.sql", "groups.sql", "group-users.sql", "messages.sql", "message-thread.sql", "notificationIds.sql", "reminders.sql", "statuses.sql", "calls.sql", "tags.sql", "visibility.sql", "alter_tables.sql"}

	// Iterate over each SQL file
	for _, file := range files {
		// if strings.HasSuffix(file.Name(), ".sql") {
		// Read the contents of the SQL file
		content, err := os.ReadFile(filepath.Join(modelsDir, file))
		if err != nil {
			log.Printf("Error reading file %s: %v\n", file, err)
			continue
		}

		// Execute the SQL statements in the file
		queries := strings.Split(string(content), ";")
		_, err1 := services.Db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
		if err1 != nil {
			log.Printf("Error installing extension %s: %v\n", file, err1)
		}

		for _, query := range queries {
			query = strings.TrimSpace(query)
			if query != "" {
				_, err := services.Db.Exec(query)
				if err != nil {
					log.Printf("Error executing query in file %s: %v\n", file, err)
				}
			}
		}

		// log.Printf("Tables created successfully from file %s\n", file)
	}
}

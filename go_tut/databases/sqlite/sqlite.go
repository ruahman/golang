package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 driver
)

func main() {
	// Open the SQLite database file. If the file does not exist, it will be created.
	// Replace "example.db" with the path to your SQLite database file.
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	fmt.Println("Connected to SQLite database")
}

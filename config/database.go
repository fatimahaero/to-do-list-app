package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "to-do-list-app.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableUser := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username STRING NOT NULL,
		password STRING NOT NULL
	);`

	createTableTask := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title STRING NOT NULL,
		parent_id INTEGER NULL,
		status BOOLEAN DEFAULT 0,
		is_delete BOOLEAN DEFAULT 0,
		user_id INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		deleted_at DATETIME NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (parent_id) REFERENCES tasks(id) ON DELETE CASCADE
	);`

	// Eksekusi query untuk membuat tabel
	_, err = db.Exec(createTableUser)
	if err != nil {
		log.Fatal("Error creating users table:", err)
	}

	_, err = db.Exec(createTableTask)
	if err != nil {
		log.Fatal("Error creating tasks table:", err)
	}

	log.Println("Database initialized successfully!")
	return db
}

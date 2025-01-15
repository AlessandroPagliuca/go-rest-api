package db

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/sqlite" // Driver SQLite
)

var DB *sql.DB

// InitDB inizializza la connessione al database SQLite.
func InitDB() {
	var err error
	// Connessione al database SQLite
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Verifica che il database sia raggiungibile
	if err := DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Impostazione delle connessioni
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Creazione delle tabelle
	createTables()
}

// createTables crea le tabelle necessarie nel database SQLite.
func createTables() {
	// Query per la creazione della tabella "events"
	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT NOT NULL,
      description TEXT NOT NULL,
      location TEXT NOT NULL,
      dateTime DATETIME NOT NULL,
      user_id INTEGER
    )
    `

	// Esegui la query
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Error creating events table: %v", err)
	}

	log.Println("Events table created or already exists.")
}

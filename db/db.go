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
	//query creazione table utenti
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatalf("Error creating events table: %v", err)
	}
	log.Println("Events table created or already exists.")

	// Query per la creazione della tabella "events"
	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES user(id)
    )
    `

	// Esegui la query
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Error creating events table: %v", err)
	}

	log.Println("Events table created or already exists.")

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		log.Fatalf("Error creating events table: %v", err)
	}

	log.Println("Registrations table created or already exists.")
}

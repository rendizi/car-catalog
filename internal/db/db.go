package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() {
	//sleep for starting postgresql
	time.Sleep(3 * time.Second)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	migrationsDir := os.Getenv("MIGRATION_DIR")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to PostgreSQL!")
	createTableQuery := `
        CREATE TABLE IF NOT EXISTS cars (
			id SERIAL PRIMARY KEY,
			regNum TEXT NOT NULL UNIQUE,
			mark TEXT NOT NULL,
			model TEXT NOT NULL,
			carYear INTEGER NOT NULL ,
			ownerName TEXT NOT NULL,
			ownerSurname TEXT NOT NULL,
			ownerPatronymic TEXT
        )
`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsDir),
		dbName, driver,
	)
	if err != nil {
		log.Fatalf("Could not create migration instance: %v", err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error applying migrations: %v", err)
	}
}

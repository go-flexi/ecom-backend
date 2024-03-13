package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // postgres driver
)

type Config struct {
	UserName    string
	Password    string
	Host        string
	Port        string
	DatabseName string
}

func (c *Config) load() {
	c.UserName = "user"
	c.Password = "password"
	c.Host = "localhost"
	c.Port = "5432"
	c.DatabseName = "ecommerce"
}

func main() {
	c := Config{}
	c.load()

	// Database connection string
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.UserName,
		c.Password,
		c.Host,
		c.Port,
		c.DatabseName,
	)

	// Establish database connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	// Initialize migrate
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create migrate driver: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // file source URL
		"postgres",          // database name
		driver,
	)
	if err != nil {
		log.Fatalf("Could not start migrate instance: %v", err)
	}

	// Apply migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while applying migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}

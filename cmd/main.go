package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	// this is the way of importing the depedenies from another folder
	"github.com/RaihanurRahman2022/simple-web-server/internal/handlers"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var db *pgxpool.Pool

func main() {
	intialize()
}

func intialize() {
	// initialize the database connection
	// load the env file and the DB Config for better security
	initDB()

	//setting up the routes here
	setupRoutes()

	//Run the Server
	RunServer()
}

// Listen and serve the server on 8080 port
// we use nil in the place of handler.
// use the default HTTP request multiplexer (http.DefaultServeMux) as the handler.
// internally the handler associated with "/" this route registered with (http.DefaultServeMux).
func RunServer() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Service down!!")
	}
}

func initDB() {
	// Under the hood Load function will load the default .env file
	// set the key values using os.Setenv function
	// later we fetch those using os.Getenv function
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// later we fetch those using os.Getenv function
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Sprintf is format a string and return it
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	var dbError error

	// pgxpool can create a pool of connections managed by pgx
	// Designed for web applications or concurrent systems where multiple queries or goroutines may need DB access.
	// Automatically opens, closes, and reuses connections.

	db, dbError = pgxpool.New(context.Background(), connStr)
	if dbError != nil {
		log.Fatalf("Unable to connect to database: %v\n", dbError)
	}

	fmt.Println("Connectd to database!")

}

func setupRoutes() {
	// create a instanct of Handler Structure
	handler := handlers.Handler{DB: db}

	// Default route Handler
	// HealthCheckHandler become a method of that Handler Structure
	http.HandleFunc("/", handler.HealthCheckHandler)

	// Events route Handler ex: /events
	//EventsHandler become a method of that Handler Structure
	http.HandleFunc("/events", handler.EventsHandler)

	// Events by id route Handler ex: /events/id
	//EventsHandler become a method of that Handler Structure
	http.HandleFunc("/events/", handler.EventHandler)

}

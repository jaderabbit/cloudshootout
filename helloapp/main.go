// [START all]
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var (
	dbhost   = "localhost"
	dbport   = 5432
	dbuser   = "test"
	dbpass   = "test"
	dbname   = "test"
	psqlInfo = ""
)

func main() {

	// Set up environment variables for database
	if fromEnv := os.Getenv("DATABASE_HOST"); fromEnv != "" {
		dbhost = fromEnv
	}

	if fromEnv, err := strconv.Atoi(os.Getenv("DATABASE_PORT")); err == nil {
		dbport = fromEnv
	}

	if fromEnv := os.Getenv("DATABASE_USER"); fromEnv != "" {
		dbuser = fromEnv
	}

	if fromEnv := os.Getenv("DATABASE_PASS"); fromEnv != "" {
		dbpass = fromEnv
	}

	if fromEnv := os.Getenv("DATABASE_NAME"); fromEnv != "" {
		dbname = fromEnv
	}

	// TODO: Move the connection into main maybe?
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpass, dbname)

	initTable()

	// use PORT environment variable, or default to 8080
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// register hello function to handle all requests
	server := http.NewServeMux()
	server.HandleFunc("/", hello)

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)

	// Seed random
	rand.Seed(time.Now().UnixNano())
}

func initTable() {

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	// Create table if not exist
	sqlStatement := `CREATE TABLE IF NOT EXISTS tests (
		id SERIAL PRIMARY KEY,
		random TEXT,
		created_at timestamptz
	);`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type MyResponse struct {
	Greeting string
	Version  string
	Host     string
	ID       int
	Msg      string
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()

	// Insert random things into db
	sqlStatement := `
		INSERT INTO tests (random, created_at)
		VALUES ($1, now())
		RETURNING id, random`
	id := 0
	txt := ""
	err = db.QueryRow(sqlStatement, randSeq(10)).Scan(&id, &txt)
	if err != nil {
		panic(err)
	}
	myResponse := MyResponse{
		Greeting: "Hello, world!",
		Version:  "1.0.0",
		Host:     host,
		ID:       id,
		Msg:      txt,
	}
	bytes, err := json.Marshal(myResponse)
	fmt.Fprintf(w, "%s", bytes)
}

// [END all]

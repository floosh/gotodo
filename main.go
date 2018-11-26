package main

import (
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
  	_ "github.com/jinzhu/gorm/dialects/sqlite"
  	"github.com/gorilla/mux"
)

var db *gorm.DB

func makeHandler(fn func(http.ResponseWriter, *http.Request, *gorm.DB)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logging
		log.Printf("%s\t%s", r.Method, r.RequestURI)
		// We're dealing with JSON here
		w.Header().Set("Content-Type", "application/json")
		// Pass DB instance to handlers
		fn(w, r, db)
	}
}


func main() {

	// Openning SQlite DB
	var err error
	db, err = gorm.Open("sqlite3", "todo.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate schemas
	db.AutoMigrate(&Todo{})

	// We're using the mux multiplexer library
	router := mux.NewRouter().StrictSlash(true)
	// API subrouter
	api := router.PathPrefix("/api").Subrouter()
	// Serve static files in 'public' folder
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// Routes !
	api.HandleFunc("/todos", 			 makeHandler(TodoIndex)).Methods("GET")		// Read
	api.HandleFunc("/todos/{id:[0-9]+}", makeHandler(TodoShow)).Methods("GET")		// Read
	api.HandleFunc("/todos", 			 makeHandler(TodoCreate)).Methods("POST")	// Create
	api.HandleFunc("/todos/{id:[0-9]+}", makeHandler(TodoUpdate)).Methods("PUT")	// Update
	api.HandleFunc("/todos/{id:[0-9]+}", makeHandler(TodoDelete)).Methods("DELETE")	// Delete

	// Server listen on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}

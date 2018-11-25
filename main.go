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

	// mux router
	router := mux.NewRouter().StrictSlash(true)

	// Routes !
	router.HandleFunc("/todos", 		makeHandler(TodoIndex)).Methods("GET")
	router.HandleFunc("/todos/{id}", 	makeHandler(TodoShow)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

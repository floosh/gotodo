package main
import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
)

func TodoIndex(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    // Return all todos
    todos := Todos{}
	if(!dbError(w, db.Find(&todos))) {
		returnJson(w, &todos);
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Grab requested Id
    id := mux.Vars(r)["id"]
    // Return the todo
    todo := Todo{}
    if(!dbError(w, db.First(&todo, id))) {
    	returnJson(w, &todo);
    }
}

func returnJson(w http.ResponseWriter, s interface{}) {
	// Json parsing
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}

// Handle db errors, return true if the db raised an error
func dbError(w http.ResponseWriter, db *gorm.DB) bool {
	// Handle errors
	if(db.RecordNotFound()) {
		// Handle record not found
		http.Error(w, db.Error.Error(), http.StatusNotFound)
	} else if db.Error != nil {
		// Everything else
		http.Error(w, db.Error.Error(), http.StatusInternalServerError)
	}
	return (db.Error != nil)
}
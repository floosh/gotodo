package main
import (
	"fmt"
	"time"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
)

func TodoIndex(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    // Return all todos
    todos := Todos{}
	if !dbError(w, db.Find(&todos)) {
		returnJson(w, &todos);
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Grab requested Id
    id := mux.Vars(r)["id"]
    // Return the todo
    todo := Todo{}
    if !dbError(w, db.First(&todo, id)) {
    	returnJson(w, &todo);
    }
}

func TodoCreate(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// setup the json decoder
	decoder := json.NewDecoder(r.Body)

    // Decode the input json to fill a Todo struct
    todo := Todo{}
    if err := decoder.Decode(&todo); err != nil {
    	http.Error(w, err.Error(), http.StatusBadRequest)
    	return
    }

    // Insert the todo
    if !dbError(w, db.Create(&todo)) {
    	// Return a 201
    	w.WriteHeader(http.StatusCreated)
    }
}

func TodoUpdate(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Grab requested Id
    id := mux.Vars(r)["id"]

    // Get the todo {id}
    todo := Todo{}
    if dbError(w, db.First(&todo, id)) {
    	return
    }
    
	// setup the json decoder
	decoder := json.NewDecoder(r.Body)
    
    // Decode the input json 
    updateData := make(map[string]interface{})
    if err := decoder.Decode(&updateData); err != nil {
    	http.Error(w, err.Error(), http.StatusBadRequest)
    	return
    }
    fmt.Printf("%T", updateData["due"])

    // Quick fix, "due" needs to be a date object
    var err error
    updateData["due"] , err = time.Parse(
        time.RFC3339,
     	updateData["due"].(string))

    // Handle bad date format
    if err != nil {
    	http.Error(w, err.Error(), http.StatusBadRequest)
    	return
    }

    // Save updated todo
    if !dbError(w, db.Model(&todo).Omit("ID").Updates(updateData)) {
    	// Return a 204
    	w.WriteHeader(http.StatusNoContent)
    }
}

func TodoDelete(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Grab requested Id
    id := mux.Vars(r)["id"]

    // Get the todo {id}
    todo := Todo{}
    if dbError(w, db.First(&todo, id)) {
    	return
    }

    // Delete the todo
    if !dbError(w, db.Delete(&todo)) {
    	// Return a 204
    	w.WriteHeader(http.StatusNoContent)
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
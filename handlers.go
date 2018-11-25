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
	if err := db.Find(&todos).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	returnJson(w, &todos);
}

func TodoShow(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Grab requested Id
    id := mux.Vars(r)["id"]
    // Return the todo
    todo := Todo{}
	if err := db.First(&todo, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	returnJson(w, &todo);
}

func returnJson(w http.ResponseWriter, s interface{}) {
	// Json parsing
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}
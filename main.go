package main

import (
	"net/http"
	"html/template"
	"log"
	"github.com/jinzhu/gorm"
  	_ "github.com/jinzhu/gorm/dialects/sqlite"

)

var templates = template.Must(template.ParseFiles("index.html", "todo.html"))
var db *gorm.DB

func viewHandler(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{}
	err := db.Find(&todos).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "index", todos)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	db.Create(&Todo{Title:title})
	http.Redirect(w, r, "/", http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, t []Todo) {
	err := templates.ExecuteTemplate(w, tmpl+".html", t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	http.HandleFunc("/",viewHandler)
	http.HandleFunc("/save/",saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

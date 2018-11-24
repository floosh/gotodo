package main

//import "github.com/jinzhu/gorm"

// Todo type
type Todo struct {
	ID 			uint 		`gorm:"primary_key"`
	Title, Desc string
}



// func (t* Todo) save() error {
// 	// Handle persistence
// 	return nil
// }
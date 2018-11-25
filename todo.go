package main

import "time"

// Todo type
type Todo struct {
	ID			uint 		`json:"id", gorm:"primary_key"`
    Title		string		`json:"title"`
    Description	string		`json:"description"`
    Completed	bool		`json:"completed"`
    Due 		time.Time 	`json:"due"`
}

type Todos []Todo
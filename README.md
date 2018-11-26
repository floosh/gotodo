# gotodo
A sample todo app, built in Go.

* Create and manage your tasks
* Set a status and a due date
* Access them with a REST API

## Implementation details
### Back-End
Built in Go using libraries below:

- [Gorm](https://github.com/jinzhu/gorm) - handle persistence with sqlite3
- [Mux](https://github.com/gorilla/mux) - HTTP request multiplexer

### Front-End
Build in HTML & Javascript, with:
- [JQuery](https://jquery.com/) - well-known library for DOM manipulation
- [Handlebars](http://handlebarsjs.com/) - HTML templates
- [Moment](http://momentjs.com/) - Time handling, parsing and rendering
- [Datepicker](https://github.com/fengyuanchen/datepicker) - Hassle-free date picking
- [Fontawesome](http://fontawesome.com/) - Fancy icons

## API schema
### Routes
```
GET     /api/todos        Get all the todos
GET     /api/todos/{id}   Get the todo {id}
POST    /api/todos        Create a todo (see todo object structure)
PUT     /api/todos/{id}   Update the todo {id}
DELETE  /api/todos/{id}   Delete the todo {id}
```
### Objects structure
Fields of a todo object:
```
id            uint
title         string
description   string
status        string
due           datetime  (format RFC3339 "2002-10-02T15:00:00Z") 
```

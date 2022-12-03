package main

import (
	"log"
	"net/http"

	r "myserver/router"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	ID       uint
}

func main() {

	// turn in mux to http

	http.Handle("/", r.Execute())

	// final servicing
	log.Print("Server Running!")
	http.ListenAndServe("localhost:8000", r.Execute())

	// hei, wait. it's not final section?

}

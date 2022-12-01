package main

import (
	"log"
	"net/http"

	r "myserver/router"
)

func main() {

	// turn in mux to http

	http.Handle("/", r.Execute())

	// final servicing
	log.Print("Server Running!")
	http.ListenAndServe("localhost:8000", r.Execute())

	// hei, wait. it's not final section?

}

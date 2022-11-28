package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// day 7 pt.1 Implementasikan materi Routing pada halaman Home, Add My Project, Detail Project & Contact Me.

	// routing

	//add mux router
	r := mux.NewRouter()

	//add func to handle each routing to work correctly
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	r.PathPrefix("/node_modules/").Handler(http.StripPrefix("/node_modules/", http.FileServer(http.Dir("./node_modules"))))
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/project", ProjectHandler).Methods("GET")
	r.HandleFunc("/contact", ContactHandler).Methods("GET")

	// turn in mux to http

	http.Handle("/", r)

	// final servicing
	fmt.Print("Server berjalan di jalan yang benar.")
	http.ListenAndServe("localhost:8000", r)

	// hei, wait. it's not final section?

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	t, e := template.ParseFiles("index.html")

	if e != nil {
		w.Write([]byte("Message : " + e.Error()))
		return
	}

	t.Execute(w, nil)

}

func ProjectHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	t, e := template.ParseFiles("view/add-project.html")

	if e != nil {
		w.Write([]byte("Message : " + e.Error()))
		return
	}

	t.Execute(w, nil)

}

func ContactHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	t, e := template.ParseFiles("view/contact.html")

	if e != nil {
		w.Write([]byte("Message : " + e.Error()))
		return
	}

	t.Execute(w, nil)

}

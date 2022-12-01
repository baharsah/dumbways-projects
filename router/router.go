package router

import (
	"net/http"

	"github.com/gorilla/mux"

	cont "myserver/controller/contact"
	homedir "myserver/controller/home"
	pro "myserver/controller/project"
)

func Execute() http.Handler {

	// day 7 pt.1 Implementasikan materi Routing pada halaman Home, Add My Project, Detail Project & Contact Me.

	// routing

	//add mux router
	r := mux.NewRouter()

	//add func to handle each routing to work correctly
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	r.PathPrefix("/node_modules/").Handler(http.StripPrefix("/node_modules/", http.FileServer(http.Dir("./node_modules"))))
	r.HandleFunc("/", homedir.HomeCtrl).Methods("GET")
	r.HandleFunc("/project", pro.ProjectCtrl).Methods("GET")
	r.HandleFunc("/contact", cont.ContactCtrl).Methods("GET")

	// day 7 pt.2 Kemudian buatlah fungsi Add My Project
	// dengan menggunakan method POST untuk mendapatkan data inputan (file input image diabaikan),
	// kemudian tampilkan kedalam Console

	r.HandleFunc("/project", pro.ProjectCtrl).Methods("POST")

	return r
}

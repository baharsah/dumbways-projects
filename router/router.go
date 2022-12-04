package router

import (
	"myserver/controller/contact"
	"myserver/controller/login"

	// "myserver/controller/project"
	"myserver/controller/home"
	"net/http"

	"github.com/gorilla/mux"
)

func Execute() http.Handler {

	// day 7 pt.1 Implementasikan materi Routing pada halaman Home, Add My Project, Detail Project & Contact Me.

	// routing

	//add mux router
	r := mux.NewRouter()

	//add func to handle each routing to work correctly
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	r.PathPrefix("/node_modules/").Handler(http.StripPrefix("/node_modules/", http.FileServer(http.Dir("./node_modules"))))
	r.HandleFunc("/", home.HomeCtrl).Methods("GET")
	// r.HandleFunc("/project", project.ProjectCtrl).Methods("GET")
	r.HandleFunc("/contact", contact.ContactCtrl).Methods("GET")

	// day 7 pt.2 Kemudian buatlah fungsi Add My Project
	// dengan menggunakan method POST untuk mendapatkan data inputan (file input image diabaikan),
	// kemudian tampilkan kedalam Console
	r.HandleFunc("/v1/register", login.RegisterAct)

	return r
}

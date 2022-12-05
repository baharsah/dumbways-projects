package project

import (
	"log"
	"myserver/model/project"
	"net/http"
	"strconv"
	"time"

	"myserver/middleware/webfs"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// struktur data project

// penampang data project

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		log.Println("Test Input")
		var sessStore = sessions.NewCookieStore([]byte("SESS_ID"))
		session, _ := sessStore.Get(r, "SESS_ID")

		r.ParseMultipartForm(1024)
		//debug
		log.Println("dari form")
		log.Println(r.FormValue("startdate"))
		postStartDate, errr := time.Parse("2006-01-02", r.FormValue("startdate"))
		postEndDate, _ := time.Parse("2006-01-02", r.FormValue("enddate"))
		log.Println(errr)
		//debug
		log.Println("setelah form")
		log.Println(postEndDate.String())
		//creating logic for file uploading
		a, b, c := r.FormFile("file")
		filename, _ := webfs.WebFSHandler(a, b, c)

		// binding to struct

		var createProject = project.Project{
			Name:        r.FormValue("name"),
			StartDate:   postStartDate,
			EndDate:     postEndDate,
			OwnerID:     session.Values["Id"].(int),
			Description: r.FormValue("desc"),
			Tech:        r.Form["tech"],
			FileDir:     filename,
		}

		//model
		err := project.CreateProject(createProject)
		log.Println(err)
		//redirecting

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])

		log.Println("Sampai Sini")

		project.DeleteProject(id)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		log.Println("Test Input")
		var sessStore = sessions.NewCookieStore([]byte("SESS_ID"))
		session, _ := sessStore.Get(r, "SESS_ID")

		r.ParseMultipartForm(1024)
		//debug
		log.Println("dari form")
		log.Println(r.FormValue("startdate"))
		postStartDate, errr := time.Parse("2006-01-02", r.FormValue("startdate"))
		postEndDate, _ := time.Parse("2006-01-02", r.FormValue("enddate"))
		log.Println(errr)
		//debug
		log.Println("setelah form")
		log.Println(postEndDate.String())
		//creating logic for file uploading
		a, b, c := r.FormFile("file")
		filename, _ := webfs.WebFSHandler(a, b, c)

		// binding to struct
		id, _ := strconv.Atoi(r.FormValue("pid"))
		var createProject = project.Project{
			Name:        r.FormValue("name"),
			StartDate:   postStartDate,
			EndDate:     postEndDate,
			OwnerID:     session.Values["Id"].(int),
			Description: r.FormValue("desc"),
			Tech:        r.Form["tech"],
			FileDir:     filename,
			ProjectID:   id,
		}

		//model
		err := project.UpdateProject(createProject)
		log.Println(err)
		//redirecting

		http.Redirect(w, r, "/", http.StatusMovedPermanently)

	}

}

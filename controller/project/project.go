package project

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Project struct {
	name        string
	startDate   string
	endDate     string
	description string
	tech        []string
}

func ProjectCtrl(w http.ResponseWriter, r *http.Request) {

	var DetailProject []Project

	if r.Method == http.MethodPost {
		err := r.ParseForm()

		if err != nil {

			fmt.Println(err)

		} else {

			// projectdata["name"] = r.FormValue("name")
			// projectdata["startDate"] = r.FormValue("startDate")
			// projectdata["endDate"] = r.FormValue("endDate")
			// projectdata["tech"] = r.Form["tech"]
			// projectdata["desc"] = r.FormValue("desc")

			var name = r.FormValue("name")
			var startDate = r.FormValue("startDate")
			var endDate = r.FormValue("endDate")
			var tech = r.Form["tech"]
			var description = r.FormValue("desc")

			var collectors = []Project{

				{

					name:        name,
					startDate:   startDate,
					endDate:     endDate,
					tech:        tech,
					description: description,
				},
			}

			append(DetailProject, collectors)

			w.Header().Set("Content-type", "text/html; charset=utf-8")
			t, e := template.ParseFiles("view/add-project.html")

			if e != nil {
				w.Write([]byte("Message : " + e.Error()))
				return
			}

			er := t.Execute(w, nil)

			if er == nil {
				log.Print(er)
			}

		}

	} else {
		w.Header().Set("Content-type", "text/html; charset=utf-8")
		t, e := template.ParseFiles("view/add-project.html")

		if e != nil {
			w.Write([]byte("Message : " + e.Error()))
			return
		}

		t.Execute(w, nil)
	}

}

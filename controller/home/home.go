package home

import (
	"log"
	dataproject "myserver/controller/project"
	"net/http"
	"text/template"

	u "myserver/model/user"
)

func HomeCtrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	t, e := template.ParseFiles("index.html")

	if e != nil {
		w.Write([]byte("Message : " + e.Error()))
		return
	}
	log.Println(dataproject.ProjectData)
	t.Execute(w, map[string]interface{}{
		"User":    u.UserData,
		"Project": dataproject.ProjectData,
	})

}

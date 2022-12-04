package home

import (
	"log"
	dataproject "myserver/controller/project"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

func HomeCtrl(w http.ResponseWriter, r *http.Request) {

	var sessStore = sessions.NewCookieStore([]byte("SESS_ID"))
	session, _ := sessStore.Get(r, "SESS_ID")

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	t, e := template.ParseFiles("index.html")

	if e != nil {
		w.Write([]byte("Message : " + e.Error()))
		return
	}
	fm := session.Flashes("Logstatus")

	var flashes []string
	if len(fm) > 0 {
		session.Save(r, w)
		// Initiate a strings slice to return messages.
		for _, fl := range fm {
			// Add message to the slice.
			flashes = append(flashes, fl.(string))
		}
	}
	log.Println(dataproject.ProjectData)
	t.Execute(w, map[string]interface{}{
		"Flashes": flashes,
		"Project": dataproject.ProjectData,
	})

}

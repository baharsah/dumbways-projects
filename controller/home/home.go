package home

import (
	"myserver/model/project"
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
	var projectCollector project.Project

	// rest, _ := strconv.Atoi(session.Values["Id"].(string))
	if session.Values["IsLogin"] != nil {

		projectCollector = project.Project{OwnerID: session.Values["Id"].(int)}

	} else {
		projectCollector = project.Project{OwnerID: 0}

	}
	// log.Println(dataproject.ProjectData)
	t.Execute(w, map[string]interface{}{
		"UserSess": map[string]interface{}{
			"IsLogin":  session.Values["IsLogin"],
			"Username": session.Values["Username"],
		},
		"Flashes": flashes,
		"Project": project.SelectProject(projectCollector),
	})

}

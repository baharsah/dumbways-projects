package contact

import (
	"html/template"
	"net/http"
)

func ContactCtrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	t, e := template.ParseFiles("view/contact.html")

	if e != nil {
		w.Write([]byte("Message : " + e.Error()))
		return
	}

	t.Execute(w, nil)

}

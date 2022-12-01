package home

import (
	"net/http"
	"text/template"
)

func HomeCtrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	t, e := template.ParseFiles("index.html")

	if e != nil {
		w.Write([]byte("Message : " + e.Error()))
		return
	}

	t.Execute(w, nil)

}

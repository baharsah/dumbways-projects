package register

import (
	"context"
	"html/template"
	"log"
	"myserver/conn"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

func RegisterAcct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	t, _ := template.ParseFiles("view/register.html")

	t.Execute(w, nil)

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		name := r.PostForm.Get("name")
		// email := r.PostForm.Get("email")

		password := r.PostForm.Get("password")
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

		_, err = conn.Conn.Exec(context.Background(), "INSERT INTO tb_user(username, password) VALUES ($1,$2,$3)", name, passwordHash)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("message : " + err.Error()))
			return
		}

		var store = sessions.NewCookieStore([]byte("SESSION_ID"))
		session, _ := store.Get(r, "SESSION_ID")

		session.AddFlash("succesfull register", "message")

		session.Save(r, w)

		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}

}

func Register(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	log.Println(r.FormValue("username"))

}

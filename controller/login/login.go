package login

import (
	"log"
	user "myserver/model/user"
	"net/http"

	"github.com/gorilla/sessions"
)

func RegisterAct(w http.ResponseWriter, r *http.Request) {

	var sessStore = sessions.NewCookieStore([]byte("SESS_ID"))
	session, _ := sessStore.Get(r, "SESS_ID")

	if r.Method == http.MethodPost {
		r.ParseForm()
		var userRegisterCollector = user.User{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		_, registerErrorCollector := user.RegisterUser(userRegisterCollector)
		if registerErrorCollector == nil {
			log.Println("Pendaftaran Berhasil!")
			session.AddFlash("Pendaftaran Berhasil! Silahkan login untuk melanjutkan.", "Logstatus")
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusMovedPermanently)

		} else {
			log.Println("Ada Error Nich!")
			log.Println(registerErrorCollector)
		}

	}
}

func LoginAcct(w http.ResponseWriter, r *http.Request) {
	var sessStore = sessions.NewCookieStore([]byte("SESS_ID"))
	session, _ := sessStore.Get(r, "SESS_ID")
	if r.Method == http.MethodPost {

		r.ParseForm()
		_, espec, id, err := user.LoginUser(user.User{Username: r.FormValue("username"), Password: r.FormValue("password")})

		if err != nil {
			log.Println(id, espec, err)
		} else {
			var sessdata = user.UserSession{ID: id, IsLogin: true}
			session.Values["IsLogin"] = sessdata.IsLogin
			session.Values["Username"] = r.FormValue("username")
			session.Values["Id"] = sessdata.ID
			session.Options.MaxAge = 10800 // 3 hours
			session.AddFlash("Login Berhasil", "Logstatus")
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}

	}
}

func LogoutAcct(w http.ResponseWriter, r *http.Request) {

	var store = sessions.NewCookieStore([]byte("SESS_ID"))
	session, _ := store.Get(r, "SESS_ID")
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)

}

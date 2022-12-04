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

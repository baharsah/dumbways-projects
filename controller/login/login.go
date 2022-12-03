package login

import (
	user "myserver/model/user"
	"net/http"
)

func LoginAcct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		if !user.UserLoginVerify(r.FormValue("username"), r.FormValue("password")) {
			// redirect to home if failure and create flashdata failure
		} else {
			user.UserData.IsLogin = true

		}
	}
}

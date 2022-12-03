package user

import (
	"context"
	"log"
	dbconn "myserver/conn"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	IsLogin  bool
	Username string
	Password string
	ID       uint
}

var UserData User

func UserLoginVerify(username string, password string) bool {
	var status bool

	dbconn.DatabaseConnect()

	user := User{}

	err := dbconn.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE username=$1", username).Scan(
		&user.ID, &user.Username, &user.Password,
	)

	if err != nil {
		log.Println(err)
		status = false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err)
		status = false
	} else {
		status = true
	}

	return status

}

func RegisterSession(user string) bool {

	return true

}

func RegisterUser(user User) {

}

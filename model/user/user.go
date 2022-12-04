package user

import (
	"context"
	"log"
	"myserver/conn"
	dbconn "myserver/conn"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
	ID       uint
}

type UserSession struct {
	IsLogin bool
	ID      int
}

type Flasher struct {
	flashdata string
}

type ErrorSpec struct {
	cause string
}

// var UserData User

func RegisterUser(user User) (bool, error) {

	var status bool
	var password = user.Password
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	conn.DatabaseConnect() // kehed sia maneh nil pointer
	_, err := dbconn.Conn.Exec(context.Background(), "INSERT INTO users(username , password) VALUES ($1, $2)", user.Username, passwordHash)
	if err != nil {
		status = false
		// log.Println("Err Here")
		log.Println(err)
		return status, err
	} else {
		status = true
		return status, err
	}

}

func LoginUser(user User) (bool, ErrorSpec, int, error) {
	// var status bool
	var userverif User
	conn.DatabaseConnect()
	err := conn.Conn.QueryRow(context.Background(), "SELECT id , username ,  password from users where username=$1", user.Username).Scan(&userverif.ID, &userverif.Username, &userverif.Password)
	if err != nil {

		return false, ErrorSpec{cause: "Username Tidak Ada"}, 0, err

	} else {
		err = bcrypt.CompareHashAndPassword([]byte(userverif.Password), []byte(user.Password))
		if err != nil {
			return false, ErrorSpec{cause: "Password Salah"}, 0, err
		} else {

			return true, ErrorSpec{}, int(userverif.ID), err
		}
	}

}

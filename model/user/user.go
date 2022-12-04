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
	ID      rune
}

type Flasher struct {
	flashdata string
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

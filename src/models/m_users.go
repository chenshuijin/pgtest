package models
import (
	"conf"
    "gopkg.in/pg.v4"
)

type User struct {
	Id, Name, Type string
}

func GetUsers() ([]User, error) {
	db := pg.Connect(conf.GetDbConf())
	defer db.Close()
	// Select all users.
	var users []User
	err := db.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, err
}

func CreateUser(user *User) error {
	db := pg.Connect(conf.GetDbConf())
	defer db.Close()
	err := db.Create(user)
	return err
}

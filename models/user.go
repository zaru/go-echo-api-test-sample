package user

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type (
	UserModel interface {
		GetDetail(id string) User
		All() []User
	}

	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	UsersModel struct {
		users []User
	}
)

func NewUsersModel() *UsersModel {
	return &UsersModel{}
}

func (u *UsersModel) GetDetail(id string) User {
	db, err := sqlx.Connect("postgres", "user=pg-user password=password dbname=sample_db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	user := User{}
	db.Select(&user, "SELECT * FROM users limit 1")
	return user
}

func (u *UsersModel) All() []User {
	u.users = append(u.users, User{
		ID:   100,
		Name: "zaru",
	})
	return u.users
}

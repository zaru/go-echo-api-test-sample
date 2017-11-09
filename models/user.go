package user

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type (
	UserModelImpl interface {
		FindByID(id string) User
		FindAll() []User
	}

	User struct {
		ID   int    `json:"id" db:"id"`
		Name string `json:"name" db:"name"`
	}

	UserModel struct {
		db *sqlx.DB
	}
)

func NewUserModel(db *sqlx.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (u *UserModel) FindByID(id string) User {
	user := User{}
	u.db.Get(&user, "SELECT * FROM users where id = $1 limit 1", id)

	return user
}

func (u *UserModel) FindAll() []User {
	users := []User{}
	u.db.Select(&users, "SELECT * FROM users order by id asc")
	return users
}

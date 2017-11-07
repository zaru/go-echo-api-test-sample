package user

type (
	UserModel interface {
		GetDetail(id string) User
		All() []User
	}

	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

func GetDetail(id string) User {
	return User{
		ID:   100,
		Name: "zaru",
	}
}

func All() []User {
	var users []User

	users = append(users, User{
		ID:   100,
		Name: "zaru",
	})
	return users
}

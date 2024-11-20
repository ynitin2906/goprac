package solidprinciples

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func (user *User) GetUserData() {
	fmt.Println(user.FirstName + " " + user.LastName)
}

type UseRepo struct{}

func (repo *UseRepo) SaveUser(user User) {
	fmt.Println(user, "User saved")
}

func RunSingleResponsibility() {
	user := User{FirstName: "Nitin", LastName: "Yadav"}
	user.GetUserData()

	repo := UseRepo{}
	repo.SaveUser(user)
}

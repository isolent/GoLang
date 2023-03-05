package models

type User struct {
	name     string
	Password string
}

func (u *User) UserInit(L string, P string) {
	u.name = L
	u.Password = P
}
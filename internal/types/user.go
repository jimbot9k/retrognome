package types

type User struct {
	ID       int
	Email    string
	Password string
	Salt     string
}

func (user *User) IsEmptyUser() bool {
	return user.ID == 0 && user.Email == "" && user.Password == "" && user.Salt == ""
}

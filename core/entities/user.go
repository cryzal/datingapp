package entities

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       string
	Email    string
	Password string
	Token    string
}

func (u *User) HashPassword() {
	password := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
func (u *User) SetToken(token string) {
	u.Token = token
}

type TokenData struct {
	ID    string
	Email string
}

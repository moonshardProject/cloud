package user

import (
	"github.com/dchest/uniuri"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID      string
	Name        string
	LastName    string
	PhoneNumber string
	Password    string
	Email       string
}

//New will return a New User
func New(
	name string,
	lastName string,
	phoneNumber string,
	email string,
	password_plain string,
) *User {
	thisUser := &User{}
	thisUser.UserID = uniuri.NewLen(12)
	thisUser.Email = email
	thisUser.LastName = lastName
	thisUser.Password = HashPassword(password_plain)
}

func Find() error {}

func Remove() error {}

func Save() error {}

// helpers

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

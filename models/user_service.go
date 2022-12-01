package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"lenslocked.com/hash"
)

// UserService interface

// UserService is a set of methods to manipulate and
// work with the user model
type UserService interface {
	// Authenticate will verify the provided email address and
	// password are correct. If they are correct the user
	// corresponding to that email will be returned. Otherwise
	// you will receive either:
	// ErrNotFound, ErrPasswordIncorrect, or another error if
	// something goes wrong
	Authenticate(email, password string) (*User, error)
	UserDB
}

// UserService implementation

type userService struct {
	UserDB
}

func NewUserService(db *gorm.DB) UserService {
	ug := &userGorm{db}
	hmac := hash.NewHMAC(hmacSecretKey)
	uv := newUserValidator(ug, hmac)
	return &userService{
		UserDB: uv,
	}
}

var _ UserService = &userService{}

// Authenticate can be used to authenticate a user with the
// provided email and password.
// If the email address provided is invalid, this will return
// nil, ErrNotFound
// If the password provided is invalid, this will return
// nil, ErrPasswordIncorrect
// If the email and password are both valid, this will return
// user, nil
// Otherwise if another error is encountered this will return
// nil, error
func (us *userService) Authenticate(email, password string) (*User, error) {
	user, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password+userPwPepper))
	switch err {
	case nil:
		return user, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return nil, ErrPasswordIncorrect
	default:
		return nil, err
	}
}

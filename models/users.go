package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
	"lenslocked.com/hash"
	"lenslocked.com/rand"
)

var (
	// ErrNotFound is returned when a resource cannot be found
	// in the database.
	ErrNotFound = errors.New("models: resource not found")
	// ErrInvalidID is returned when an invalid id is passed to methods
	// like Delete.
	ErrInvalidID = errors.New("models: ID provided was invalid")
	// ErrInvalidPassword is returned when an invalid password
	// is used when attempting to authenticate a user.
	ErrInvalidPassword = errors.New("models: incorrect password provided")

	userPwPepper = "secret-secret-secret"
	key          = "secret-hmac-key"
)

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;unique_index"`
}

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	hmac := hash.NewHMAC(key)
	return &UserService{
		db:   db,
		hmac: hmac,
	}, nil
}

type UserService struct {
	db   *gorm.DB
	hmac hash.HMAC
}

func (us *UserService) Close() error {
	return us.db.Close()
}

// Create will create the provided user and back-fill data
// like the ID, CreatedAt and UpdatedAt fields.
func (us *UserService) Create(user *User) error {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password+userPwPepper), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""

	if user.Remember == "" {
		user.Remember, err = rand.RememberToken()
		if err != nil {
			return err
		}
	}
	user.RememberHash = us.hmac.Hash(user.Remember)
	err = us.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// Authenticate can be used to authenticate a user with the
// provided email and password.
// If the email address provided is invalid, this will return
// nil, ErrNotFound
// If the password provided is invalid, this will return
// nil, ErrInvalidPassword
// If the email and password are both valid, this will return
// user, nil
// Otherwise if another error is encountered this will return
// nil, error
func (us *UserService) Authenticate(email, password string) (*User, error) {
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
		return nil, ErrInvalidPassword
	default:
		return nil, err
	}
}

// ByID will look up a user with the provided ID.
// If the user is found, we will return a nil error
// If the user is not found, we will return ErrNotFound
// If there is another error we will return an error with
// more information about was wrong. This may not be
// an error generated by the models package.
//
// As a general rule, any error but ErrNotFound should
// probably result in a 500 error.
func (us *UserService) ByID(id uint) (*User, error) {
	return getUserByCondition(us.db, "id = ?", id)
}

// ByEmail will look up a user with the provided email.
func (us *UserService) ByEmail(email string) (*User, error) {
	return getUserByCondition(us.db, "email = ?", email)
}

// ByRemember looks up a user with the given remember token
// and will return that user. The method will handle hashing the token.
func (us *UserService) ByRemember(token string) (*User, error) {
	return getUserByCondition(us.db, "remember_hash = ?", us.hmac.Hash(token))
}

func getUserByCondition(db *gorm.DB, query string, args ...interface{}) (*User, error) {
	var user User
	db = db.Where(query, args)
	if err := first(db, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

// Update will update the provided user with all the data
// in the provided user object.
// if a user with such id absent, creates a new entry
func (us *UserService) Update(user *User) error {

	// Note: we don't handle changing the password yet (set new password hash)

	if user.Remember != "" {
		user.RememberHash = us.hmac.Hash(user.Remember)
	}
	return us.db.Save(user).Error
}

// UpdateRememberToken updates remember token in the user object,
// not touching other attributes to avoid potential race conditions
func (us *UserService) UpdateRememberToken(user *User) error {

	// Note: we don't handle changing the password yet (set new password hash)

	if user.Remember != "" {
		user.RememberHash = us.hmac.Hash(user.Remember)
	}
	return us.db.Model(user).Update("remember_hash", us.hmac.Hash(user.Remember)).Error
}

// Delete deletes a user by id. Actually just sets deleted_at
// to a non-null value, so the entry is recoverable
// GORM deletes all data if given 0 as an id, so returning
// an ErrInvalidID for id == 0
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{
		Model: gorm.Model{ID: id},
	}
	return us.db.Delete(&user).Error
}

// AutoMigrate will automatically migrate the
// users table
func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}

// DestructiveReset drops the user table and resets it
func (us *UserService) DestructiveReset() error {
	err := us.db.DropTableIfExists(&User{}).Error
	if err != nil {
		return err
	}
	return us.AutoMigrate()
}

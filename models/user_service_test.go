package models

import (
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	userName = "aru"
	password = ""
	dbname   = "lenslocked_test"
)

var us UserService
var s *Services

func init() {

	// :(
	//var err error
	//s, err := NewServices(
	//	WithGorm(dbCfg.Dialect(), dbCfg.ConnectionInfo()),
	//	WithLogMode(!cfg.IsProd()),
	//	WithUser(cfg.Pepper, cfg.HMACKey),
	//	WithGallery(),
	//	WithImage(),
	//)

	//if err != nil {
	//	panic(err)
	//}
	//us = s.User
	//s.DestructiveReset()
}

func TestUserModel(t *testing.T) {

	user := User{
		Name:     "anton",
		Email:    "rum@ya.ru",
		Password: "qwerty123",
	}

	us.Create(&User{
		Name:     "kek",
		Email:    "kek@cheburek.com",
		Password: "qwerty123",
	})

	us.Create(&User{
		Name:     "cheburek",
		Email:    "test@lol.com",
		Password: "qwerty123",
	})

	// Create a user
	err := us.Create(&user)
	if err != nil {
		t.Error(err)
	}

	// Query for a user
	u, err := us.ByID(user.ID)
	if err != nil {
		t.Error(err)
	}
	if u.Name != user.Name || u.Email != user.Email {
		t.Errorf("Got %+v, want %+v", u, user)
	}

	// Query for a user by email
	_, err = us.ByEmail("rum@ya.ru")
	if err != nil {
		t.Error(err)
	}

	// Query for a user, check Error
	_, err = us.ByID(500)
	if err != ErrNotFound {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {

	user := User{
		Name:     "user for update",
		Email:    "user@user.com",
		Password: "qwerty123",
	}

	err := us.Create(&user)
	if err != nil {
		t.Error(err)
	}

	user.Name = "updated"
	user.Email = "updated@updated.com"

	err = us.Update(&user)
	if err != nil {
		t.Error(err)
	}

	found, err := us.ByID(user.ID)
	if err != nil {
		t.Error(err)
	}
	if found.Name != user.Name {
		t.Errorf("name was not updated, want \"%s\", got \"%s\"", user.Name, found.Name)
	}
	if found.Email != user.Email {
		t.Errorf("email was not updated, want \"%s\", got \"%s\"", user.Email, found.Email)
	}

}

func TestDelete(t *testing.T) {
	user := User{
		Name:     "delete",
		Email:    "delete@delete.com",
		Password: "qwerty123",
	}
	err := us.Create(&user)
	if err != nil {
		t.Error(err)
	}
	err = us.Delete(user.ID)
	if err != nil {
		t.Error(err)
	}
	_, err = us.ByID(user.ID)
	if err != ErrNotFound {
		t.Error("user was not deleted")
	}
}

func TestUserService_Authenticate(t *testing.T) {
	err := us.Create(&User{
		Name:     "name_auth",
		Email:    "email@auth.com",
		Password: "authauthauth",
	})
	if err != nil {
		t.Error(err)
	}

	_, err = us.Authenticate("email@auth.com", "authauthauth")
	if err != nil {
		t.Error(err)
	}
}

func TestUserService_ByRemember(t *testing.T) {
	err := us.Create(&User{
		Name:     "do you",
		Email:    "remember@remember.com",
		Password: "turututu",
		Remember: "rememberrememberrememberrememberrememberrememberrememberrememberremember",
	})
	if err != nil {
		t.Error(err)
	}

	u, err := us.ByRemember("rememberrememberrememberrememberrememberrememberrememberrememberremember")
	if err != nil {
		t.Error(err)
	}
	if u.Remember != "" {
		t.Error("Remember token was saved, but should not be")
	}
}

func TestDestruct(t *testing.T) {
	err := s.DestructiveReset()
	if err != nil {
		t.Error(err)
	}
}

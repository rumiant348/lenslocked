package controllers

import (
	"fmt"
	"lenslocked.com/models/users"
	"lenslocked.com/rand"
	"lenslocked.com/views"
	"net/http"
)

func NewUsers(us users.UserService) *Users {
	return &Users{
		NewView:   views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:        us,
	}
}

type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        users.UserService
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// New is used to render a new form where a user can
// create a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create is used to process the signup form when a user
// tries to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	var form SignupForm
	if err := ParseForm(r, &form); err != nil {
		vd.SetAlert(err)
		u.NewView.Render(w, vd)
		return
	}

	user := &users.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	err := u.us.Create(user)
	if err != nil {
		vd.SetAlert(err)
		u.NewView.Render(w, vd)
		return
	}
	err = u.signIn(w, user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

// Login is used to process the login form when a user
// tries to log in as an existing user (via email & pw)
//
// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	form := LoginForm{}
	if err := ParseForm(r, &form); err != nil {
		vd.SetAlert(err)
		u.LoginView.Render(w, vd)
		return
	}
	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case users.ErrNotFound:
			vd.AlertError("No user exists with that email address")
		default:
			vd.SetAlert(err)
		}
		u.LoginView.Render(w, vd)
		return
	}
	err = u.signIn(w, user)
	if err != nil {
		vd.SetAlert(err)
		u.LoginView.Render(w, vd)
		return
	}
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

func (u *Users) signIn(w http.ResponseWriter, user *users.User) error {
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
		err = u.us.Update(user)
	}

	cookie := http.Cookie{
		Name:  "remember_token",
		Value: user.Remember,
		// disables scripting usage, set to false for react
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return nil
}

// CookieTest is used to display cookies on the current user
func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("remember_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := u.us.ByRemember(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, user)
}
